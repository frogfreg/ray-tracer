package main

import (
	"fmt"
	"ray-tracer/rays"
	tpv "ray-tracer/tuplespointsvectors"
)

func main() {

	origin := tpv.NewPoint(0, 0, 5)
	direction := tpv.NewVector(0, 0, 1)

	r := rays.NewRay(origin, direction)

	s := rays.NewSphere()

	xs := rays.Intersect(s, r)

	fmt.Println(xs)
}
