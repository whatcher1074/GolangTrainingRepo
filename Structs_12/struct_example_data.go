package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	fmt.Println("Getting started")
	sa1 := secretAgent{
		person: person{
			first: "Wendell",
			last:  "Hatcher",
			age:   47,
		},
		ltk: true,
	}

	p1 := person{}
	p2 := person{
		first: "Ashley",
		last:  "Williams",
		age:   38,
	}

	p1.first = "Wendell"
	p1.last = "Hatcher"
	p1.age = 47

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)

	bar("Bleh bleh bleh")
	q := woo("Running")
	fmt.Println(q)
	x, y := mouse("Dell", "Hatcher")
	fmt.Println(x, y)
	foo(1, 2, 3, 4, 5, 6, 7, 8, 10)
	sa1.speak()
}

func bar(s string) {
	fmt.Println(s)
}

func woo(q string) string {
	return fmt.Sprint("Hello back to a string ", q)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false
	return a, b

}

func foo(x ...int) {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func (s secretAgent) speak() {
	fmt.Println("Testing new data", s.first, s.last)
}
