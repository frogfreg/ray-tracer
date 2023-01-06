package canvas

import (
	"fmt"
	"math"
	tpv "ray-tracer/tuplespointsvectors"
	"strings"
)

type Canvas struct {
	width, height int

	Pixels [][]tpv.Tuple
}

func NewCanvas(width, height int) Canvas {

	canvas := Canvas{}

	canvas.width = width
	canvas.height = height

	pixels := make([][]tpv.Tuple, height)

	for i := range pixels {

		pixels[i] = make([]tpv.Tuple, width)

	}

	canvas.Pixels = pixels

	return canvas

}

func (c *Canvas) WritePixel(row, column int, color tpv.Tuple) {

	c.Pixels[column][row] = color

}

func (c *Canvas) ToPPM() string {

	var builder strings.Builder
	line := ""

	dimensions := fmt.Sprintf("%v %v\n", c.width, c.height)

	builder.WriteString("P3\n")
	builder.WriteString(dimensions)
	builder.WriteString("255\n")

	for _, row := range c.Pixels {

		for _, col := range row {

			if len(line)+12 >= 70 {
				builder.WriteString("\n")
				line = ""
			}

			var r int
			var g int
			var b int

			r = int(math.Floor(col.X * 255))
			g = int(math.Floor(col.Y * 255))
			b = int(math.Floor(col.Z * 255))

			if col.X <= 0 {
				r = 0
			}
			if col.X >= 1 {
				r = 255
			}
			if col.Y <= 0 {
				g = 0
			}
			if col.Y >= 1 {
				g = 255
			}
			if col.Z <= 0 {
				b = 0
			}
			if col.Z >= 1 {
				b = 255
			}

			if len(line) != 0 {
				line += " "
			}

			line += fmt.Sprintf("%v %v %v", r, g, b)
			builder.WriteString(line)

		}

		builder.WriteString("\n")

	}

	return builder.String()

}
