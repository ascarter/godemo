package main

import (
	"fmt"
	"time"
)

func produce(msg string, delay time.Duration, c chan<- string) {
	time.Sleep(delay); c <- msg
}

func main() {
	c1, c2 := make(chan string), make(chan string)
	go produce("one", time.Second*1, c1)
	go produce("two", time.Second*2, c2)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1: fmt.Println("received", msg1)
		case msg2 := <-c2: fmt.Println("received", msg2)
		case <-time.After(time.Second): fmt.Println("timeout")
		}
	}
}
