package main

import "fmt"

func main() {

	sum := []int{1, 2, 3, 4, 5, 8, 9, 78, 6, 4}
	results := foo(sum...)
	fmt.Println("getting results from foo: ", results)
	sums := []int{2, 4, 5, 7, 8}
	fmt.Println(bar(sums))
}

func foo(x ...int) int {
	add := 0
	for _, v := range x {
		fmt.Println("the value of v is:", v)
		add += v
	}

	return add

}

func bar(x []int) int {
	add := 0
	for _, v := range x {
		fmt.Println("the value of v is:", v)
		add += v
	}

	return add

}
