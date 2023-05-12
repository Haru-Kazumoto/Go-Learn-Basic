package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist("User") {
		fmt.Println("User is not assignable on this pages!")
	} else {
		fmt.Println("Welcome back!")
	}
}

func main() {
	registerUser("User", func(name string) bool {
		return name == "User"
	})	
}