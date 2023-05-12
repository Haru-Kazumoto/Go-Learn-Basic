package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Sofi"
	names[1] = "Sahwatul"
	names[2] = "Amirah"

	fmt.Println(names)

	var number = [3]int{
		1,
		2,
		3,
	}

	var changed = number[1]
	number[0] = number[1]
	number[1] = number[2]
	number[2] = changed

	fmt.Println(number)
	fmt.Println(changed)
}
