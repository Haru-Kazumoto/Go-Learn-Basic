package main

import (
	"errors"
	"fmt"
)

func divider(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Math cannot dividing 0!")
	} else {
		return nilai / pembagi, nil
	}
}

func main(){
	result, err := divider(10,0)
	fmt.Println(result, err.Error())
}
