package main

import (
	"fmt"
	"ray-tracer/rays"
	tpv "ray-tracer/tuplespointsvectors"
)

func main() {

	origin := tpv.NewPoint(2, 3, 4)
	direction := tpv.NewVector(1, 0, 0)

	r := rays.NewRay(origin, direction)

	fmt.Println(r.Position(2.5))

}
