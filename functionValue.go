package main

import "fmt"

func main() {
	hello := sayHi
	fmt.Println(hello("RRRAULLL"))
}

func sayHi(word string) string {
	return "Hello" + word
}