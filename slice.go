package main

import "fmt"

func main() {
	var month = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"Oktober",
		"November",
		"December",
	}

	var slice1 = month[5:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// There is a good way to create a slice array, like this
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ziaurrahman"
	newSlice[1] = "Athaya"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//This is a way to copy a slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//Note:
	//The different between array and slice is
	//when u create like this way 
	//variable := []int{12345} //This is an slice, because there is no value in bracket []
	//variable := [...] or [5]int{123456} //This is an array, cause there is a value in bracket.
}
