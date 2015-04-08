package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() { 
		for i := 0; i < 10; i++ {
			messages <- fmt.Sprintf("ping%d", i)
		}
		close(messages)
	}()
	
	for m := range messages {
		fmt.Println(m)
	}
}
