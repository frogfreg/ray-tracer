package main

import (
	"os"
	"ray-tracer/canvas"
	tpv "ray-tracer/tuplespointsvectors"
)

func main() {
	// p := tpv.NewPoint(0, 1, 0)

	canv := canvas.NewCanvas(5, 5)

	color := tpv.Newrgb(12, 242, 93)

	canv.WritePixel(2, 2, color)

	// fmt.Println(canv)

	os.WriteFile("./images/clock.ppm", []byte(canv.ToPPM()), 0666)

}
