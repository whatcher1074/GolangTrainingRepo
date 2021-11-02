package main

import "fmt"

func main() {

	fmt.Println("Starting a new program")
	y := foo(5)
	fmt.Println("Printing reutrned valus for foo: ", y(10))
	d := 8
	q := 6
	g := bar(q, d)
	fmt.Println("Printing reutrned valus for bar: ", g)

}

func foo(f int) func(int) int {
	return func(a int) int {
		return bar(f, a)
	}
}

func bar(a int, z int) int {
	b := a + 85
	c := z + 92
	return b + c

}
