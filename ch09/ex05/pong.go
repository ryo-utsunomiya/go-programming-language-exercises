package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ping := make(chan struct{})
	pong := make(chan struct{})
	cnt := 0

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				pong <- <-ping
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				ping <- <-pong
				cnt++
			}
		}
	}()

	ping <- struct{}{}

	<-ctx.Done()
	fmt.Println(cnt) // 1967254
}
