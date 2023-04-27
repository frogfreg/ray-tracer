package main

import (
	"fmt"
	"ray-tracer/rays"
	tpv "ray-tracer/tuplespointsvectors"
)

func main() {

	origin := tpv.NewPoint(0, 0, -5)
	direction := tpv.NewVector(0, 0, 1)

	r := rays.NewRay(origin, direction)
	s := rays.NewSphere()

	// i1 := *rays.NewIntersection(1, s)
	// i2 := *rays.NewIntersection(2, s)

	xs := s.Intersect(r)

	if xs[0].Object == xs[1].Object {
		fmt.Println("the object is the same sphere")
	}

	fmt.Printf("%#v\n", xs)
}
