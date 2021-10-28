package main

import (
	"fmt"
	"runtime"
)

var y int = 56

const s = 42
const q = "string"

func main() {
	// short identifier declares the var initalization then it can be reassigned.
	h := make([]int, 17)
	h = append(h, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(h)
	j := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, j := range j {
		fmt.Println(i, j)

	}
	fmt.Println(j)

	m := map[string]int{
		"Testing": 1,
		"Second":  2,
	}

	m["Dell"] = 47
	m["Asjley"] = 38
	delete(m, "Testing")

	fmt.Println(m)
	for k, v := range m {

		fmt.Println(k, v)
	}

	x := 42
	a := 26
	type bleh int

	var b bleh
	fmt.Println("running first data set")
	n, _ := fmt.Println("Getting Data back as well as return typ information")
	fmt.Println(n)
	fmt.Println(x)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println(s, q)

	a = int(b)
	// new value assignment

	fmt.Println(x)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	z := "New data and some stuff"
	fmt.Println(z)

	testing()

}

func testing() {

	fmt.Println("Testing we can print infered y value data outside of scope", y)
}
