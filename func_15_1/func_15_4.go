package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	people := person{}
	people.first = "Dell"
	people.last = "Hatcher"
	people.age = 35
	people.speak()
}

func (a person) speak() {
	fmt.Println("My name is", a.first, " ", a.last, " and my age is: ", a.age)

}
