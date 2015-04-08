package main

import "fmt"

func test() {
	var message = "Hello"
	defer func() {
		fmt.Println("Message was", message)
	}()
	fmt.Println("Welcome to test")
}

func main() {
	test()
}
