package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() { 
		for i := 0; i < 20; i++ {
			messages <- fmt.Sprintf("ping%d", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(messages)
	}()
	
	for m := range messages {
		fmt.Println(m)
	}
}
