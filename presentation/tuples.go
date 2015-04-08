package main

import (
	"fmt"
	"strings"
)

func parse(name string) (first, last string) {
	parts := strings.Split(name, " ")
	first, last = parts[0], parts[1]
	return
}

func main() {
	names := []string{"Tim Cook", "Bill Gates", "Elon Musk"}
	for _, name := range names {
		first, last := parse(name)
		fmt.Println("Hello", first, last)
	}
}
