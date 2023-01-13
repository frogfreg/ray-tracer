package main

import (
	"fmt"
	"ray-tracer/matrix"
	_ "ray-tracer/tuplespointsvectors"

	"github.com/shopspring/decimal"
)

func main() {

	// a := matrix.NewMatrix(2, 2)

	// a[0][0] = decimal.NewFromFloat(1)
	// a[0][1] = decimal.NewFromFloat(5)
	// a[1][0] = decimal.NewFromFloat(-3)
	// a[1][1] = decimal.NewFromFloat(2)

	// for _, row := range a {
	// 	fmt.Println(row)
	// }

	// fmt.Printf("The determinant is: %v", a.Determinant())

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

	a[0][0] = decimal.NewFromFloat(3)
	a[0][1] = decimal.NewFromFloat(-9)
	a[0][2] = decimal.NewFromFloat(7)
	a[0][3] = decimal.NewFromFloat(3)
	a[1][0] = decimal.NewFromFloat(3)
	a[1][1] = decimal.NewFromFloat(-8)
	a[1][2] = decimal.NewFromFloat(2)
	a[1][3] = decimal.NewFromFloat(-9)
	a[2][0] = decimal.NewFromFloat(-4)
	a[2][1] = decimal.NewFromFloat(4)
	a[2][2] = decimal.NewFromFloat(4)
	a[2][3] = decimal.NewFromFloat(1)
	a[3][0] = decimal.NewFromFloat(-6)
	a[3][1] = decimal.NewFromFloat(5)
	a[3][2] = decimal.NewFromFloat(-1)
	a[3][3] = decimal.NewFromFloat(1)

	for _, row := range a {
		fmt.Println(row)
	}

	b := matrix.NewMatrix(4, 4)

	b[0][0] = decimal.NewFromFloat(8)
	b[0][1] = decimal.NewFromFloat(2)
	b[0][2] = decimal.NewFromFloat(2)
	b[0][3] = decimal.NewFromFloat(2)
	b[1][0] = decimal.NewFromFloat(3)
	b[1][1] = decimal.NewFromFloat(-1)
	b[1][2] = decimal.NewFromFloat(7)
	b[1][3] = decimal.NewFromFloat(0)
	b[2][0] = decimal.NewFromFloat(7)
	b[2][1] = decimal.NewFromFloat(0)
	b[2][2] = decimal.NewFromFloat(5)
	b[2][3] = decimal.NewFromFloat(4)
	b[3][0] = decimal.NewFromFloat(6)
	b[3][1] = decimal.NewFromFloat(-2)
	b[3][2] = decimal.NewFromFloat(0)
	b[3][3] = decimal.NewFromFloat(5)

	fmt.Println()

	for _, row := range b {
		fmt.Println(row)
	}

	c := matrix.Multiply(a, b)

	fmt.Println()

	for _, row := range c {
		fmt.Println(row)
	}

	d := matrix.Multiply(c, b.Inverse())

	fmt.Println()

	for _, row := range d {
		fmt.Println(row)
	}
	fmt.Println()

	fmt.Println("Are equal?", matrix.AreEqual(a, d))

	// for _, row := range a {
	// 	fmt.Println(row)
	// }

	// fmt.Println()

	// for _, row := range b {
	// 	fmt.Println(row)
	// }

	// fmt.Println()

	// c := matrix.Multiply(a, b)

	// for _, row := range c {
	// 	fmt.Println(row)
	// }

	// fmt.Println("Are equal? ", matrix.AreEqual(a, matrix.Multiply(c, b.Inverse())))

	// fmt.Printf("The determinant is: %v\n", a.Determinant())
	// fmt.Printf("Is invertible?: %v\n\n", a.IsInvertible())

	// fmt.Println("determinant of a: ", a.Determinant())
	// fmt.Println("cofactor of a: ", matrix.Cofactor(a, 2, 3))
	// fmt.Println("b[3,2]: ", b[3][2])
	// fmt.Println("cofactor(A,3,2): ", matrix.Cofactor(a, 3, 2))
	// fmt.Println("b[2,3]: ", b[2][3])

	// fmt.Println()

	// fmt.Println("Are equal?", matrix.AreEqual(c, a))
	// b := matrix.NewMatrix(4, 4)

	// fmt.Println()

	// for _, row := range b {
	// 	fmt.Println(row)
	// }

	// fmt.Printf("multiplied a * b\n\n")

	// c := matrix.Multiply(a, b)

	// fmt.Println("c:")
	// fmt.Println()

	// for _, row := range c {
	// 	fmt.Println(row)
	// }

	// d := matrix.Multiply(c, b.Inverse())

	// fmt.Println("d == a? ", matrix.AreEqual(a, d))

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
