package canvas

import (
	"errors"
	"fmt"
	"math"
	tpv "ray-tracer/tuplespointsvectors"
	"strings"
)

type Canvas struct {
	Width, Height int

	Pixels [][]tpv.Tuple
}

func (c Canvas) PixelAt(x, y int) (tpv.Tuple, error) {
	if x > c.Width-1 || x < 0 {
		return tpv.Tuple{}, fmt.Errorf("index out of range")
	}
	if y > c.Height-1 || y < 0 {
		return tpv.Tuple{}, fmt.Errorf("index out of range")
	}
	return c.Pixels[y][x], nil
}

func NewCanvas(width, height int) Canvas {

	canvas := Canvas{}

	canvas.Width = width
	canvas.Height = height

	pixels := make([][]tpv.Tuple, height)

	for i := range pixels {

		pixels[i] = make([]tpv.Tuple, width)

	}

	canvas.Pixels = pixels

	return canvas

}

func (c *Canvas) WritePixel(x, y int, color tpv.Tuple) error {

	if y > c.Height-1 || y < 0 {
		return errors.New("index out of range")
	}
	if x > c.Width-1 || x < 0 {
		return errors.New("index out of range")
	}

	c.Pixels[y][x] = color
	return nil

}

func (c *Canvas) ToPPM() string {

	var builder strings.Builder
	line := ""

	dimensions := fmt.Sprintf("%v %v\n", c.Width, c.Height)

	builder.WriteString("P3\n")
	builder.WriteString(dimensions)
	builder.WriteString("255\n")

	for _, row := range c.Pixels {
		for _, col := range row {

			if len(line)+12 >= 70 {
				builder.WriteString(line + "\n")
				line = ""
			}

			var r int
			var g int
			var b int

			r = int(math.Ceil(col.X * 255))
			g = int(math.Ceil(col.Y * 255))
			b = int(math.Ceil(col.Z * 255))

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

		}

		builder.WriteString(line)
		builder.WriteString("\n")
		line = ""
	}

	return builder.String()
}
