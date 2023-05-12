package main

import "fmt"

func main() {
	var name string
	name = "Ziaurrahman"

	//Or you can do this way
	var age = 17

	//And if you want to expilicit the data type of variable you can do this way
	var dob uint16 = 2006

	// actually the var keyword to initialize the variable is not mandatory,
	// but must immediately fill in the value of the variable and use := to provide its value.
	// But := is only used for the first initialize, and then you can initialize without : first.
	city := "Bekasi"

	fmt.Println("Hi! my name is", name, ", and im", age, "years old. I was born in", dob, "in", city)

	// There is a way to create a multiple variable in one step, you can use this way
	var (
		createdOn = "Bekasi"
		onDate    = "11/05/2023"
	)

	fmt.Println("Created on : ", createdOn)
	fmt.Println("On date    : ", onDate)

}
