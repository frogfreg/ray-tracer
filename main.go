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

	a[0][0] = 3
	a[0][1] = -9
	a[0][2] = 7
	a[0][3] = 3
	a[1][0] = 3
	a[1][1] = -8
	a[1][2] = 2
	a[1][3] = -9
	a[2][0] = -4
	a[2][1] = 4
	a[2][2] = 4
	a[2][3] = 1
	a[3][0] = -6
	a[3][1] = 5
	a[3][2] = -1
	a[3][3] = 1

	b := matrix.NewMatrix(4, 4)

	b[0][0] = 8
	b[0][1] = 2
	b[0][2] = 2
	b[0][3] = 2
	b[1][0] = 3
	b[1][1] = -1
	b[1][2] = 7
	b[1][3] = 0
	b[2][0] = 7
	b[2][1] = 0
	b[2][2] = 5
	b[2][3] = 4
	b[3][0] = 6
	b[3][1] = -2
	b[3][2] = 0
	b[3][3] = 5

	for _, row := range a {
		fmt.Println(row)
	}

	fmt.Println()

	for _, row := range b {
		fmt.Println(row)
	}

	c := matrix.Multiply(a, b)

	fmt.Println("c:")
	fmt.Println()

	for _, row := range c {
		fmt.Println(row)
	}

	d := matrix.Multiply(c, b.Inverse())

	fmt.Println("d == a? ", matrix.AreEqual(a, d))

	// fmt.Println("d:")
	// fmt.Println()

	// for _, row := range d {
	// 	fmt.Println(row)
	// }

	// fmt.Printf("The determinant is: %v\n", a.Determinant())
	// fmt.Printf("Is invertible?: %v\n\n", a.IsInvertible())

	// fmt.Println("determinant of a: ", a.Determinant())
	// fmt.Println("cofactor of a: ", matrix.Cofactor(a, 2, 3))
	// fmt.Println("b[3,2]: ", b[3][2])
	// fmt.Println("cofactor(A,3,2): ", matrix.Cofactor(a, 3, 2))
	// fmt.Println("b[2,3]: ", b[2][3])

	// for _, row := range a {
	// 	fmt.Println(row)
	// }

	// c := matrix.Multiply(a, b)

	// fmt.Println()

	// for _, row := range c {
	// 	fmt.Println(row)
	// }
}
