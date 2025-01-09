package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Circum() float64
}

type Square struct {
	length float64
}

type Circle struct {
	radius float64
}

func (s Square) Area() float64 {
	return s.length * s.length
}

func (s Square) Circum() float64 {
	return s.length * 4
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) Circum() float64 {
	return c.radius * math.Pi * 2
}

func printshapeInfo(s Shape) {
	fmt.Println("Area of Square is :", s, s.Area())
	fmt.Println("Circumference of Square is :", s, s.Circum())
	fmt.Println("Area of Circle is :", s, s.Area())
	fmt.Println("Circumference of Circle is :", s, s.Circum())
}

func main() {
	Shape := []Shape{
		Square{length: 15.2},
		Circle{radius: 7.5},
	}

	for _, v := range Shape {
		printshapeInfo(v)
		fmt.Println("----")
	}
}
