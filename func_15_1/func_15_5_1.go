package main

import "fmt"

func main() {

	c := 98
	fmt.Println("Some new stuff before my anonmouse func call")
	// secret blank func call
	func(i int) {
		fmt.Println("Calling this in secret then leaving")
		v := i + 55
		fmt.Println("Print out the results of v as well", v)
	}(c)

}
