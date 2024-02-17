package main

import (
	"fmt"
	"time"
)

var count int = 0

func ping(c chan string) {
	for {
		msg := <-c
		if msg == "ping" {
			fmt.Println("ping")
			time.Sleep(time.Second * 1)
			c <- "pong"
		}
	}
}

func pong(c chan string) {
	for {
		msg := <-c
		if msg == "pong" {
			fmt.Println("pong")
			time.Sleep(time.Second * 1)
			c <- "ping"
		}
	}
}

func startPing(c chan string) {
	c <- "ping"
}

func main() {
	var c chan string = make(chan string)

	go ping(c)
	go pong(c)

	go startPing(c)

	var entrada string
	fmt.Scanln(&entrada)
}
