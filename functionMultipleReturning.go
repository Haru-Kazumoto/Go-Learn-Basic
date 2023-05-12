package main

import "fmt"

func main() {

	// Like for range method, you can use only one key and the other key who ignored change to _
	// example : firstName, _ := getFullName()
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}

func getFullName() (string, string) {
	return "Ziaurrahman", "Athaya"
}
