package main

import (
	"math"
	"os"
	"ray-tracer/canvas"
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

	p := projectile{tpv.Point(0, 1, 0), tpv.ScMult(tpv.Normalized(tpv.Vector(1, 1.8, 0)), 11.25)}
	e := environment{tpv.Vector(0, -0.1, 0), tpv.Vector(-0.01, 0, 0)}

	c := canvas.NewCanvas(900, 550)
	color := tpv.Color(0, 1, 0)

	c.WritePixel(int(math.Ceil(p.position.X)), c.Height-int(math.Ceil(p.position.Y)), color)

	for {

		if p.position.Y <= 0 {
			break
		}

		// fmt.Printf("Current position %v\n", p.position)

		p = tick(e, p)

		x := int(math.Ceil(p.position.X))
		y := int(math.Ceil(p.position.Y))

		c.WritePixel(x, c.Height-y, color)
	}

	header := c.ToPPM()

	os.WriteFile("./projectile.ppm", []byte(header), 0666)

}
