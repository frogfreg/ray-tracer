package main

import (
	"os"
	"ray-tracer/canvas"
	tpv "ray-tracer/tuplespointsvectors"
)

func main() {

	canvas := canvas.NewCanvas(5, 3)

	c1 := tpv.Newrgb(1.5, 0, 0)
	c2 := tpv.Newrgb(0, 0.5, 0)
	c3 := tpv.Newrgb(-0.5, 0, 1)

	canvas.WritePixel(0, 0, c1)
	canvas.WritePixel(2, 1, c2)
	canvas.WritePixel(4, 2, c3)

	header := canvas.ToPPM()

	os.WriteFile("./temp.txt", []byte(header), 0666)
}
