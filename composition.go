package main

import "fmt"

type Name struct {
	First string
	Last  string
}

func (n *Name) String() string {
	return fmt.Sprintf("%s %s", n.First, n.Last)
}

func main() {
	name := Name{"Tim", "Cook"}
	fmt.Println("Hello", name.String())
}
