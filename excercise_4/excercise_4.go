package main

import "fmt"

type myOwnType int

var x myOwnType

func main() {

	fmt.Println("Initialize fmt print call")
	fmt.Println("Print value of x before assigning a number  ", x)
	fmt.Printf("Print out the type  %T ", x)
	//assign x a value of 42
	x := 42
	fmt.Println("\n X has been assigned a new value of 42:", x)

}
