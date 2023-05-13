package main

import "fmt"

func random() interface{} {
	return "zero"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)

	fmt.Println(resultString)

	//But if you type assertioning the wrong data type, it will throw panic
}