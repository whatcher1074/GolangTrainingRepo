package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	raduis float64
}

type square struct {
	lenght float64
}

func main() {
	c := circle{
		raduis: 12.345,
	}
	sqau := square{
		lenght: 15,
	}

	info("circle area:", c)
	info("sqaure area:", sqau)
}

func (c circle) area() float64 {
	return math.Pi * c.raduis * c.raduis
}

func (s square) area() float64 {

	return s.lenght * s.lenght

}

func info(t string, s shape) {
	x := s.area()
	fmt.Println(t, x)

}
