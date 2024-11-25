package main

import "fmt"

func main() {
	var firstName string = "Riccardo"
	var lastName string = "Greco"

	var dayOfBirth byte = 31
	var monthOfBirth byte = 07
	var yearOfBirth int16 = 2008

	var numberOfSiblings byte = 2

	var heightInMeters float64 = 1.86

	var zodiacSign rune = '\u264c'

	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
