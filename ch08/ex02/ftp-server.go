package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const CRLF = "\r\n"

func main() {
	var (
		port = flag.Int("port", 2222, "port")
	)
	flag.Parse()
	address := fmt.Sprintf("localhost:%d", *port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go server(conn)
	}

}

func server(c net.Conn) {
	mustWrite(c, "220 Ready.")
	b := bufio.NewReader(c)

	for {
		cmd, _, err := b.ReadLine()
		if err != nil {
			continue
		}
		processCommand(c, string(cmd))
	}
}

func processCommand(c net.Conn, cmd string) {
	switch strings.ToLower(cmd) {
	case "pwd":
		pwd, err := os.Getwd()
		if err != nil {
			return
		}
		mustWrite(c, fmt.Sprintf("257 %s", pwd))
	case "close":
		c.Close()
	default:
		mustWrite(c, "502 Command not implemented")
	}
}

func mustWrite(w io.Writer, text string) {
	fmt.Println(text)
	_, err := w.Write([]byte(text + CRLF))
	if err != nil {
		log.Fatal(err)
	}
}
