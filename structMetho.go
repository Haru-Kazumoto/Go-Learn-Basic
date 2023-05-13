package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func  halo(person Person) {
	fmt.Println("Hello my name is ",person.Name,",Im",person.Age,"years old.")
}

func main() {
	var objPerson Person
	objPerson.Name = "Haru"
	objPerson.Age = 17

	halo(objPerson)
}