package main

import (
	"fmt"
	tpv "ray-tracer/tuplespointsvectors"
)

type projectile struct {
	position, velocity tpv.Tuple
}

type environment struct {
	gravity, wind tpv.Tuple
}

func tick(e environment, p projectile) projectile {

	newProj := projectile{}

	newProj.position = tpv.Add(p.position, p.velocity)

	newProj.velocity = tpv.Add(tpv.Add(p.velocity, e.gravity), e.wind)

	return newProj

}

func main() {

	p := projectile{tpv.NewPoint(0, 1, 0), tpv.Normalized(tpv.NewVector(1, 1, 0))}
	e := environment{tpv.NewVector(0, -0.1, 0), tpv.NewVector(-0.01, 0, 0)}

	for {

		if p.position.Y <= 0 {
			break
		}

		fmt.Printf("Current position %v\n", p.position)

		p = tick(e, p)

	}

}
