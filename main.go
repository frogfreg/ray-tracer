package main

import (
	"math"
	"os"

	"ray-tracer/canvas"
	"ray-tracer/matrix"
	tpv "ray-tracer/tuplespointsvectors"
)

func main() {
	canv := canvas.New(250, 250)

	twelvePoint := tpv.Point(0, .7, 0)

	halfCanvasWidth := float64(canv.Width / 2)
	halfCanvasHeight := float64(canv.Height / 2)

	sliceOfPi := (math.Pi * 2) / 12

	for i := range 12 {
		invRotationMat := matrix.RotationZ(float64(i) * (sliceOfPi)).Inverse()

		p := matrix.TupleMultiply(twelvePoint, invRotationMat)
		x := int(halfCanvasWidth * (p.X + 1))
		y := int(halfCanvasHeight * (1 - p.Y))

		if err := canv.WritePixel(x, y, tpv.Color(0, 1, 0)); err != nil {
			panic(err)
		}

		if err := os.WriteFile("./clock.ppm", []byte(canv.ToPPM()), 0o666); err != nil {
			panic(err)
		}
	}
}
