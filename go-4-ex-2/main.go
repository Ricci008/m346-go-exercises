package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Println(computeHypotenuse(1, 4))
	fmt.Println(computeHypotenuse(5, 18))
	fmt.Println(computeHypotenuse(9, 15))
}
