package main

import "fmt"

func main() {
	total := sumAll(10, 123, 14, 234, 21341, 2341)
	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}
