package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan struct{})
	pong := make(chan struct{})
	cnt := 0

	go func() {
		for {
			cnt++
			pong <- <-ping
		}
	}()
	go func() {
		for {
			ping <- <-pong
		}
	}()

	quits := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		quits <- struct{}{}
	}()

	ping <- struct{}{}
	<-quits
	fmt.Println(cnt) // 2267670
}
