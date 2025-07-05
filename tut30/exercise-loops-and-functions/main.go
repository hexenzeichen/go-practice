package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("my sqrt is %v, theirs sqrt is %v\n", Sqrt(3), math.Sqrt(3))

}

func Sqrt(x float64) float64 {

	z := 1.0
	d := z

	for d > 0.000000001 {
		d = z*z - x
		z -= d / (2 * z)
		if d < 0 {
			d = 0 - d
		}
	}

	return z
}
