package main

import (
	"fmt"
)

func main() {

	for counter := 1; counter < 10; counter++ {
		fmt.Println("loop - [", counter, "]")
	}

	// Iterate a slice using for
	slice := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
	}

	slice = append(slice, "eight")

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//  key    value
	for index, number := range slice {
		fmt.Println("Index", index, "value - ", number)
	}

	// But if you want only get the value without using the index or the key,
	// basically in go lang if a variable is ignored it will error, then if you want only use the value
	// you can change the key to _ for initialize the key is ignore or not used.
	for _, number := range slice {
		fmt.Println("Value - ", number)
	}

	person := make(map[string]string)
	person["name"] = "Haru Kazumoto"
	person["age"] = "17"

	for key, value := range person {
		fmt.Println(key, "-", value)
	}
}
