package main

import (
	"log"
	"net"
	"time"
	"fmt"
	"encoding/json"
	"bufio"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8083")
	if err != nil {
		log.Fatal("net error：", err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("listener error", err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	name string
	messages chan string
	conn *net.Conn
}

var (
	entering = make(chan *client, 2)
	leaving  = make(chan *client, 2)
	messages = make(chan string, 4) // all incoming client messages
)

func broadcaster() {
	clients := make(map[*client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			log.Println(msg)
			for cli := range clients {
				cli.messages <- msg
			}
		case cli := <-entering:
			clients[cli] = true
			messages <- cli.name + " has arrived"
		case cli := <-leaving:
			messages <- cli.name + " has left"
			delete(clients, cli)
			close(cli.messages)
			(*cli.conn).Close()
		}
	}
}

func handleConn(conn net.Conn) {
	cli := &client{
		name:conn.RemoteAddr().String(),
		messages:make(chan string),
		conn: &conn,
	}
	log.Println(cli.name, "connected")
	go clientWriter(cli)
	entering <- cli
	hasInput := make(chan bool)
	go timeoutClose(hasInput, cli)
	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()
		data := decodeResponse([]byte(text))
		if name, ok := data["name"]; ok{
			cli.name = name
			cli.messages <- "You are " + cli.name
		}
		if msg, ok := data["msg"]; ok {
			messages <- "[" + cli.name + "]" + ": " + msg
		}
		hasInput <- true
	}
}

func decodeResponse(data []byte) map[string]string {
	res := map[string]string{}
	err := json.Unmarshal(data, &res)
	if err != nil {
		log.Println("json decode error:", err, data)
		return map[string]string{"msg":string(data)}
	}
	return res
}


func clientWriter(cli *client) {
	for msg := range cli.messages {
		_, err := fmt.Fprintln(*cli.conn, msg)
		if err != nil {
			//log.Println("tcp connection error:", err)
		}
	}
}

func timeoutClose(hasInput <-chan bool, cli *client)  {
	for {
		select {
		case <-hasInput:
			continue
		case <-time.After(300*time.Second): // 5分钟没有输入，自动关闭
			 leaving <- cli
			return
		}
	}
}