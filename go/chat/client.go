package main

import (
	"net"
	"bufio"
	"fmt"
	"os"
	"log"
	"encoding/json"
	"strings"
)

func main() {
	conn, _ := net.Dial("tcp", "localhost:8083")

	fmt.Println("What's you name?")
	name, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		log.Fatal("read from stdin error:", err)
	}
	go sendData(&conn, map[string]string{
		"name":strings.TrimSuffix(name, "\n"),
	})

	go func() {
		for {
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				log.Fatal("read from connet error:", err)
			}
			log.Print(message)
		}
	}()
	for {
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal("read from stdin error:", err)
		}
		go sendData(&conn, map[string]string{"msg": strings.TrimSuffix(text, "\n")})
	}
}

func sendData(conn *net.Conn, data interface{}) (error){
	t, err := json.Marshal(data)
	if err != nil {
		log.Println("json error:", err)
		return err
	}
	n, err := fmt.Fprint(*conn, string(t) + "\n")
	if err != nil {
		log.Println("connection error:", err)
	}
	log.Println(string(t), n)
	return err
}