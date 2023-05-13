package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var alamat = Address{
		City:     "Bekasi",
		Province: "Jawa barat",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}