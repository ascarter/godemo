package main

import "fmt"

type Name struct {
	First string
	Last  string
}

func (n *Name) Greeting() string {
	return fmt.Sprintf("Hello %s %s!", n.First, n.Last)
}

func main() {
	name := Name{"Tim", "Cook"}
	fmt.Println(name.Greeting())
}
