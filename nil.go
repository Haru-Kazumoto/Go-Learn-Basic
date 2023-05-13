package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	result := newMap("")
	if result == nil{
		fmt.Println("Data is empty!")
	} else {
		fmt.Println(result)
	}
}