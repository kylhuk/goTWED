package main

import (
	"fmt"
	"github.com/kylhuk/goTWED"
)

func main() {

	ta := []float64{0, 0, 1, 1, 2, 3, 5, 2, 0, 1, -0.1}
	tsa := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tb := []float64{0, 1, 1, 1, 2, 3, 5, 2, 0, 1, -0.1}
	tsb := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nu := float64(1.1)
	lambda := float64(2.2)
	degree := int32(1)

	dist := gotwed.TWED(ta, tsa, tb, tsb, nu, lambda, degree)

	if dist != -1 {
		fmt.Println("The calculated distance is: ", dist)
	}
}
