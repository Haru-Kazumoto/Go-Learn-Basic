package main

import "fmt"

/**
if you want to add any parameter that ignore the data type (any), you can add like this way
func print(i int, anyData interface{}){} //the parameter will be recieve any data type no matter what is it
*/
func print(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 0 {
		return "0"
	} else {
		return false
	}
}

/**
	Empty interface can be write like this
	interface{} and any
	both are same way, any is just representing interface{}

	data = map[string]any / data = map[string]interface{}

	when u want to set the data type explicitly without put data type in {}
	you could use converting like variable = data.(int64)
*/

func main() {
	//If you want to print the empty interface, you can only initialize the variable
	//interface{} type, if you initialize other data type, it will error.

	//var data int = print(2) //this variable will throw an error.
	var data interface{} = print(2)
	fmt.Println(data)
}