package main

import "fmt"

func main() {

	name := "icikiwir"

	switch name {
	case "Haru":
		fmt.Println("Helo",name)
	case "haru":
		fmt.Println("Halo",name)
	default: 
		fmt.Println("Nama lu bukan haru coy")
	}

	//You can use short statement too
	switch length := len(name); length > 5{
	case true:
		fmt.Println("Nama lu kepanjangan")
	case false:
		fmt.Println("Oke!")
	}

}