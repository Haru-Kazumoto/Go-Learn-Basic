package main

import "fmt"

type Person struct {
	Name string
}

func (person *Person) StatusPerson() {
	person.Name = "Haru"
}

func main() {
	person1 := Person{"Kazumoto"}
	person1.StatusPerson()

	fmt.Println(person1.Name)
}