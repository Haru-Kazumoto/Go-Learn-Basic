package main

import "fmt"

func main() {
	forEach()
}

func forEach() {
	for i := 0; i < 10; i++ {
		fmt.Println("Index - ", i)
	}
}
