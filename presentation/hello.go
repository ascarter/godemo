package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	s := "Xcoders"

	fmt.Printf("Hello %s. Today is %v", s, t.Format(time.Stamp))
}
