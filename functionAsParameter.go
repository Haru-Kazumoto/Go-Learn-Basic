package main

import "fmt"

func main() {
	filterWords("botak asu", spamFilter)
}

type Filter func(string) string //For initialize the function in parameter

func filterWords(name string, filter Filter) {
	filtering := filter(name)
	fmt.Println(filtering)
}

func spamFilter(name string) string{
	if name == "botak asu"{
		return "**** ***"
	} else {
		return name
	}
}