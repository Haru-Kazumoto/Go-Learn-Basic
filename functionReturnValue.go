package main

import "fmt"

func main() {
	type (
		num1 int
		num2 int
	)

	var (
		number1 num1 = 14
		number2 num2 = 34
	)

	fmt.Println(
		additional(int(number1), int(number2)),
	)
}

func additional(num1 int, num2 int) int {
	result := num1 + num2
	return result
}
