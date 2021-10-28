package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Lets get the slice type %T\n", a)
	for i, a := range a {
		fmt.Println(i, a)

	}

}
