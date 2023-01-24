package main

import (
	"fmt"
	"math"
	"os"
	"ray-tracer/canvas"
	"ray-tracer/matrix"
	tpv "ray-tracer/tuplespointsvectors"
)

func main() {

	hourPoints := []tpv.Tuple{}

	hourPoints = append(hourPoints, tpv.NewPoint(0, 1, 0))

	for i := 1; i < 12; i++ {

		transform := matrix.NewIdentityMatrix(4, 4).RotateZ(((2 * math.Pi) / 12) * float64(i))

		newP := matrix.TupleMultiply(hourPoints[0], transform)

		hourPoints = append(hourPoints, newP)

	}

	canv := canvas.NewCanvas(100, 100)
	color := tpv.Newrgb(242, 61, 76)

	for _, p := range hourPoints {

		x := p.X * -40
		y := p.Y * -40

		x += 49
		y += 49

		intX := int(math.Ceil(x))
		intY := int(math.Ceil(y))

		fmt.Printf("x: %v, y:%v\n", intX, intY)

		canv.WritePixel(intX, intY, color)
		os.WriteFile("./images/clock.ppm", []byte(canv.ToPPM()), 0666)

	}
}
