package main

import "fmt"

func main() {

	//Type is for used like interface in typescript, it could re-use it
	type (
		studentName  string
		gradeStudent int8
		NISNStudent  string
		gender       string
		hasGraduate  bool
	)

	// This is how to insert the value of type declaration.
	var (
		name     studentName  = "Haru"
		grade    gradeStudent = 11
		nisn     NISNStudent  = "12341235418236"
		Gender   gender       = "Male"
		graduate hasGraduate  = false
	)

	fmt.Println("Name 		:", name)
	fmt.Println("Grade 		:", grade)
	fmt.Println("NISN 		:", nisn)
	fmt.Println("Gender 	:", Gender)
	fmt.Println("Graduate 	:", graduate)
}
