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

	// a := matrix.NewMatrix(3, 3)

	// a[0][0] = 1
	// a[0][1] = 2
	// a[0][2] = 6
	// a[1][0] = -5
	// a[1][1] = 8
	// a[1][2] = -4
	// a[2][0] = 2
	// a[2][1] = 6
	// a[2][2] = 4

	a := matrix.NewMatrix(4, 4)

	a[0][0] = -2
	a[0][1] = -8
	a[0][2] = 3
	a[0][3] = 5
	a[1][0] = -3
	a[1][1] = 1
	a[1][2] = 7
	a[1][3] = 3
	a[2][0] = 1
	a[2][1] = 2
	a[2][2] = -9
	a[2][3] = 6
	a[3][0] = -6
	a[3][1] = 7
	a[3][2] = 7
	a[3][3] = -9

	for _, row := range a {
		fmt.Println(row)
	}

	fmt.Println()

	fmt.Printf("The determinant is: %v\n", a.Determinant())
	fmt.Printf("Is invertible?: %v\n\n", a.IsInvertible())

	for _, row := range a {
		fmt.Println(row)
	}
}
