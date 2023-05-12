package main

import "fmt"

func main() {

	//So whenever the variable is typed by const, it cannot be edited with other variable
	const firstName = "Ziaurrahman"
	const lastName = "Athaya"

	//So if you try to edit the variable like this :
	// firstname = "edited";
	//It will throw an error.

	//Basically data type const simmilar like final in java.

	fmt.Println("HI! my name is", firstName, lastName)

	// And you can you multiple variable tho
	const (
		first_Name string = "Ziaurrahman"
		last_Name         = "Athaya"
		age        int8   = 17
		dob        int16  = 2006
	)

	fmt.Println("Hi! my name is", first_Name, last_Name, ", im", age, "years old. I was born in", dob)
}
