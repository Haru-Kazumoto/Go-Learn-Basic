package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address
	address1.City = "Bekasi"
	address1.Province = "Jawa barat"
	address1.Country = "Indonesia"

	var address2 *Address = &address1

	address2.City = "Tangerang"

	fmt.Println(address1)
	fmt.Println(address2)
}
