package main

import "fmt"

var x = 42
var y = "Wendell Hatcher"
var z = true

func main() {

	fmt.Println("Just something to use import fmt call")
	s := fmt.Sprintf("%v  %v %v \n", x, y, z)
	fmt.Println(s)

}
