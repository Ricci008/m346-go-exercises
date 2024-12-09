package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints, maxPoints float64) float64 {
	return (gotPoints/maxPoints)*5 + 1
}

func main() {
	// TODO: call the function computeGrade
	fmt.Println(computeGrade(17.5, 28.0))
	fmt.Println(computeGrade(25.0, 30.0))
	fmt.Println(computeGrade(07.0, 20.0))
}
