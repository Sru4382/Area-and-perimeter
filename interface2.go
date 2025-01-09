package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	breath float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.breath
}

func printArea(s shape) {
	fmt.Printf("%v area is %v\n", s, s.area())
}

func main() {
	r := rectangle{breath: 4, length: 5}
	c := circle{radius: 5}
	printArea(r)
	printArea(c)
}
