package main

import (
	"fmt"
	"ray-tracer/matrix"
	tpv "ray-tracer/tuplespointsvectors"
)

func main() {
	transform := matrix.Shearing(0, 0, 0, 0, 0, 1)
	// fullQuarter := matrix.RotateZ(math.Pi / 2)

	// v := tpv.NewVector(2, 3, 4)

	// inv := halfQuarter.Inverse()

	p := tpv.NewPoint(2, 3, 4)

	res1 := matrix.TupleMultiply(p, transform)
	// res2 := matrix.TupleMultiply(p, fullQuarter)

	fmt.Println(p)

	fmt.Println(res1, "Is point?", res1.IsPoint(), "Is vector?", res1.IsVector())
	// fmt.Println(res2, "Is point?", res2.IsPoint(), "Is vector?", res2.IsVector())

}
