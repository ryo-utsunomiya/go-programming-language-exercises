package main

import (
	"bufio"
	"flag"
	"fmt"
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
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Print("ftp>")

	c := bufio.NewReader(conn)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		_, err := conn.Write([]byte(s.Text() + "\r\n"))
		if err != nil {
			log.Fatal(err)
		}

		b, _, err := c.ReadLine()
		if err != nil {
			break
		}
		fmt.Println(string(b))

		fmt.Print("ftp>")
	}
}
