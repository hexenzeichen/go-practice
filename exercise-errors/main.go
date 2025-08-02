package main

import (
	"fmt"
)


func Sqrt(x float64) (float64, error) {

	z := 1.0
	d := z

	for d > 0.000000001 {
		d = z*z - x
		z -= d / (2 * z)
		if d < 0 {
			d = 0 - d
		}
	}

	return z, nil
}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
