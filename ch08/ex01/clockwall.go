package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

type clock struct {
	name    string
	address string
	time    string
}

func (c *clock) String() string {
	return fmt.Sprintf("%s\t%s", c.name, c.time)
}

func (c *clock) Tick(reader io.Reader) {
	r := bufio.NewReader(reader)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		c.time = string(b)
	}
}

func main() {
	clocks := make([]*clock, 0)
	for _, arg := range os.Args[1:] {
		parts := strings.Split(arg, "=")
		c := clock{
			name:    parts[0],
			address: parts[1],
		}
		clocks = append(clocks, &c)
		go timer(&c)
	}
	for {
		time.Sleep(1 * time.Second)
		clear()
		for _, c := range clocks {
			fmt.Println(c)
		}
	}
}

func timer(c *clock) {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c.Tick(conn)
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
