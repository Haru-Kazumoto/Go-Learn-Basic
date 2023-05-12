package main

import "fmt"

func logging() {
	message := recover()
	if message != nil{
		fmt.Println("Error with message : ", message)
	}
	fmt.Println("Request function.")
}

func runApplication(error bool) {
	defer logging()
	if error {
		panic("Internal logic error!")
	}
	fmt.Println("Application is success.")
}

func main(){
	runApplication(true)
}