package main

import "fmt"

type Student struct {
	Name    string
	Address string
	NISN    string
	Age     int
}

func main() {

	//This is basic usage of struct type 
	var objStudent Student

	objStudent.Name = "Raull"
	objStudent.Address = "Mumbai"
	objStudent.Age = 19
	objStudent.NISN = "13298746238746"

	fmt.Println(objStudent)
	fmt.Println(objStudent.Name)

	//This is struct literal method
	budi := Student{
		Name: "Budi",
		Address: "Bekasi",
		NISN: "91283471982",
		Age: 18,
	}

	fmt.Println(budi)

	// Also you can do this way
	raul := Student{"Raul","Bekasi","128362187",18}

	fmt.Println(raul)
}
