package main

import (
	"flag"
	"net"
	"log"
	"io"
	"time"
)


func main() {
	var port = flag.String("port", "8000", "请输入端口")
	flag.Parse()
	println(*port)
	listener, err := net.Listen("tcp", "localhost:" + *port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle one connection at a time
	}
}
func handleConn(c net.Conn) { defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}