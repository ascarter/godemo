package main

import (
	"fmt"
	"strings"
)

func parse(location string) (city, state string) {
	parts := strings.Split(location, ",")
	city, state = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	return
}

func main() {
	locations := []string{"Seattle, WA", "San Francisco, CA", "Portland, OR"}
	for _, loc := range locations {
		city, state := parse(loc)
		fmt.Printf("Welcome to %s in %s\n", city, state)
	}
}
