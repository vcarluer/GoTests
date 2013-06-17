package main

import (
	"fmt"
	"math"
)

func main() {
	sqrt(2)
}

func sqrt(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / (2 * z))
	}

	fmt.Printf("Newton:%v, Real:%v", z, math.Sqrt(x))

	return x;
}
