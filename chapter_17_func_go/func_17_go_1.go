package main

import "fmt"

type Person struct {
	First_Name string
	Last_Name  string
	Address    string
	age        int
}

func main() {

	fmt.Println("Starting App:")
	NewPerson := Person{}

	NewPerson.First_Name = "Wendell"
	NewPerson.Last_Name = "Hatcher"
	NewPerson.age = 45

	fmt.Println(NewPerson)
	changeMe(&NewPerson)
	fmt.Println(NewPerson)

}

func changeMe(p *Person) {
	(*p).First_Name = "Chuck"

}
