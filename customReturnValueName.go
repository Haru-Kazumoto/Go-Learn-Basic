package main

import "fmt"

func main() {
	key1, key2, key3 := getInformation()
	fmt.Println(key1, key2, key3)
}

func getInformation() (firstName string, lastName string, age int) {
	firstName = "Ziaurrahman"
	lastName = "Athaya"
	age = 17

	return
}
