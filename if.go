package main

import "fmt"

func main() {
	var name = "Haru"

	if name == "Haru" {
		fmt.Println(true)
	} else if name == "haru" {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	//This is short statement if
	if length := len(name); length > 5 {
		fmt.Println("Name is too long!")
	} else {
		fmt.Println("Name is ok!")
	}
}
