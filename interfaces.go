package main

import "fmt"

type Shape interface {
	area() float64
}

type Rectangle struct{ l, w float64 }

type Square struct{ s float64 }

func (r Rectangle) area() float64 { return r.l * r.w }

func (s Square) area() float64 { return s.s * s.s }

func main() {
	shapes := []Shape{Rectangle{10, 20}, Square{5}}
	for _, s := range shapes {
		fmt.Printf("area = %f\n", s.area())
	}
}
