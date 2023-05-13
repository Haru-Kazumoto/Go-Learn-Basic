package main

import "fmt"

func random() interface{} {
	return "Hello!"
}
/**
	Type assertion is change variable or converting the type data from default is non typed and 
	give the value of what type data is required in empty interface.
*/
func main() {
	var result interface{} = random()
	// var resultString string = result.(string)

	// fmt.Println(resultString)

	//But if you type assertioning the wrong data type, it will throw panic

	//For more safe, you could use switch statement instead variable.

	switch value := result.(type){
	case string :
		fmt.Println("String type = ", value)
	case int:
		fmt.Println("Int type = ", value)
	default :
		fmt.Println("Unknown type = ", value) //The default is when the type data is not match the logic.
	}
}