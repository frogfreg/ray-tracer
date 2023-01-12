package main

import (
	"fmt"
	"ray-tracer/matrix"
)

func main() {

	// a := matrix.NewMatrix(2, 2)

	// a[0][0] = 1
	// a[0][1] = 5
	// a[1][0] = -3
	// a[1][1] = 2

	// for _, row := range a {
	// 	fmt.Println(row)
	// }

	// fmt.Printf("The determinant is: %v", matrix.SimpleDeterminant(a))

	a := matrix.NewMatrix(3, 3)

	a[0][0] = 3
	a[0][1] = 5
	a[0][2] = 0
	a[1][0] = 2
	a[1][1] = -1
	a[1][2] = -7
	a[2][0] = 6
	a[2][1] = -1
	a[2][2] = 5

	for _, row := range a {
		fmt.Println(row)
	}

	fmt.Println()

	fmt.Printf("The cofactor is: %v", matrix.Cofactor(a, 1, 0))
}
