package main

import "fmt"

func main() {
	greetings("Ziaurrahman", "Athaya")
}

func greetings(firstName string, lastName string) {
	fmt.Println("Hello!")
	fmt.Println("My name is", firstName, lastName)
}
