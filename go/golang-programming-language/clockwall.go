package main

import (
	"os"
	"strings"
	"time"
	"net"
	"log"
	"io"
)

func main() {
	for _, v := range os.Args[1:] {
		keyValue := strings.Split(v, "=")
		go connTcp(keyValue[1])
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

func connTcp(uri string) {
	conn, err := net.Dial("tcp", uri)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}