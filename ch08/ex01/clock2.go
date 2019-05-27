package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	var (
		port = flag.Int("port", 8000, "port")
	)
	flag.Parse()

	timezone := os.Getenv("TZ")
	if timezone == "" {
		timezone = "Asia/Tokyo"
	}
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalf("invalid timezon: %s", timezone)
	}
	time.Local = loc

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
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
