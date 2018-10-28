package main

import "fmt"

// exercise

func main() {

	var z,x float64
	z = 1
	x= 144

	for i := 0; i < 10; i++ {
		z = calcSquare(z,x)

		fmt.Println(z)
	}
}

func calcSquare (z, x float64) float64 {
	z -= (z*z - x) / (2*z)
	return z
}