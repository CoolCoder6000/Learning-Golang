package main

import (
	"gonum.org/v1/gonum/mat"
)

func getDet(matrix mat.Matrix) (determinant float64) {
	determinant = mat.Det(matrix)
	return
}

func main() {
	// fmt.Println("Hello World!")

	// Create a new matrix
	// a := mat.NewDense(3, 3, []float64{
	// 	1, 2, 3,
	// 	4, 5, 6,
	// 	7, 8, 9,
	// })

	// b := mat.NewDense(3, 3, []float64{
	// 	1, 0, 0,
	// 	0, 1, 0,
	// 	0, 0, 1,
	// })

	// fmt.Println(getDet(a))
	// fmt.Println(getDet(b))

	// Calculate determinant
	// det := mat.Det(a)

	// Calculate inverse
	// var inv mat.Dense
	// err := inv.Inverse(a)
}
