package main

import "fmt"

func main() {

	//Converting data type in go lang is simple, you can do this way
	var (
		nilai32 int32 = 100000
		nilai64 int64 = int64(nilai32)
		nilai8  int8  = int8(nilai32)
	)

	// However, in the int32 and int8 data types,
	// basically the limit of the numbers is not the same
	// and when int8 has reached the specified number limit,
	// it will return to the default number or minus number and so on ,
	// until it reaches the limit again.

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8) // result = -96

	// This is how to convert the string type
	var name = "Mira"
	var byteName byte = name[0]
	var byteNameToString string = string(byteName)

	fmt.Println(name)
	fmt.Println(byteNameToString)
}
