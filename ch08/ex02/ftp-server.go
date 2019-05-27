package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

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
	switch cmd {
	case "pwd":
		pwd, err := os.Getwd()
		if err != nil {
			return
		}
		mustPrint(c, pwd)
	case "close":
		c.Close()
	default:
		mustPrint(c, fmt.Sprintf("unknown command: %s", cmd))
	}
}

func mustPrint(w io.Writer, text string) {
	fmt.Println(text)
	_, err := w.Write([]byte(text + "\r\n"))
	if err != nil {
		log.Fatal(err)
	}
}
