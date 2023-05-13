package main

import "fmt"

//Interface can be re use in other function
type HasName interface {
	GetName() string
}

func sayHi(name HasName) {
	fmt.Println("Hello", name.GetName())
}

type Person struct {
	Name string
}

type User struct {
	Username string
}

//Implement 1
func (person Person) GetName() string {
	return person.Name
}

//Implement 2
func (user User) GetName() string{
	return user.Username
}

func main() {
	var budi Person
	budi.Name = "Budi"

	user := User{
		Username: "Kazumoto",
	}
	
	sayHi(budi)
	sayHi(user)

}
