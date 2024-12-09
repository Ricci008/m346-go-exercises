package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a, b, c float64) []float64 {
	discriminant := math.Pow(b, 2) - 4*a*c
	if discriminant > 0 {
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return []float64{x1, x2}
	} else if discriminant == 0 {
		x := -b / (2 * a)
		return []float64{x}
	} else {
		return []float64{}
	}

}

func main() {
	// TODO: call the function computeQuadraticFormula
	fmt.Println(computeQuadraticFormula(1, 2, 1))
	fmt.Println(computeQuadraticFormula(1, 3, 2))
	fmt.Println(computeQuadraticFormula(3, 4, 2))
}
