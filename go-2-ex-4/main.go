package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Studens struct {
		FirstName string
		LastName  string
	}
	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Studens
	}
	// TODO: declare a map of modules being attended by multiple classes
	type Module struct {
		Number  int
		Classes []Class
	}
	// TODO: output everything using fmt.Println()
	student1 := Studens{FirstName: "Riccardo", LastName: "Greco"}
	student2 := Studens{FirstName: "John", LastName: "Doe"}
	student3 := Studens{FirstName: "Jane", LastName: "Doe"}
	student4 := Studens{FirstName: "Max", LastName: "Mustermann"}
	student5 := Studens{FirstName: "Erika", LastName: "Mustermann"}
	student6 := Studens{FirstName: "Hans", LastName: "Peter"}

	class1 := Class{Students: []Studens{student1, student2, student3}}
	class2 := Class{Students: []Studens{student4, student5, student6}}

	module1 := Module{Number: 111, Classes: []Class{class1, class2}}
	module2 := Module{Number: 222, Classes: []Class{class1}}
	module3 := Module{Number: 333, Classes: []Class{class2}}

	fmt.Println("Klasse 1:", class1)
	fmt.Println("Klasse 2:", class2)

	fmt.Println("Module 111:", module1)
	fmt.Println("Module 222:", module2)
	fmt.Println("Module 333:", module3)
}
