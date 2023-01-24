package main

import (
	"fmt"
	"ray-tracer/rays"
	tpv "ray-tracer/tuplespointsvectors"
)

func main() {

	origin := tpv.NewPoint(1, 2, 3)
	direction := tpv.NewVector(4, 5, 6)

	r := rays.NewRay(origin, direction)

	fmt.Println(tpv.SameTuple(origin, r.Origin))
	fmt.Println(tpv.SameTuple(direction, r.Direction))
}
