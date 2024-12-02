package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	LastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	// TODO: embed full name and birth date information
	Name             FullName
	Birth            BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "Riccardo",
			LastName:  "Greco",
		},
		Birth: BirthDate{
			Day:   31,
			Month: 07,
			Year:  2008,
		},
		// TODO: set name and birth date information
		NumberOfSiblings: 2,        // TODO: adjust
		ZodiacSign:       '\u264c', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
