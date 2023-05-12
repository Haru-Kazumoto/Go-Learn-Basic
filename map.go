package main

import "fmt"

func main() {

	//Map is key:value, simmilar like Map<K,V> in java.
	person := map[string]string{
		"name": "Ziaurrahman",
		"age":  "17",
	}

	//You can add value like this way 
	person["grade"] = "High school"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["grade"])

	//Create new map use keyword make()
	var student = make(map[string]string)

	student["name"] = "Haru Kazumoto"
	student["email"] = "HaruKazumoto@gmail.com"
	student["age"] = "17"
	student["wrong"] = "This field will be deleted"

	fmt.Println(len(student))
	delete(student,"wrong")

	fmt.Println(student)
	fmt.Println(len(student))
}