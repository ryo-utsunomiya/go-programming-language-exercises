package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	defer c.Close()

	input := make(chan struct{})

	go func() {
		for {
			timeout := time.After(10 * time.Second)
			select {
			case <-timeout:
				c.Close()
			case <-input:
				timeout = time.After(10 * time.Second)
			}
		}
	}()

	in := bufio.NewScanner(c)
	for in.Scan() {
		input <- struct{}{}
		go echo(c, in.Text(), 1*time.Second)
	}
	close(input)
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
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
