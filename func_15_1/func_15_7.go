package main

import "fmt"

func main() {

	fmt.Println("Starting a new program")
	y := foo(5, 6)
	fmt.Println("Printing reutrned valus for foo: ", y)
	d := 8
	q := 6
	g, h := bar(q, d)
	fmt.Println("Printing reutrned valus for bar: ", g, h)

}

func foo(f int, q int) int {

	a := f + q

	return a
}

func bar(a int, z int) (int, int) {
	b := a + 85
	c := z + 92
	return b, c

}
