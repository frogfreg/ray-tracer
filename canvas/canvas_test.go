package canvas

import (
	"testing"

	tpv "ray-tracer/tuplespointsvectors"
)

func TestNewCanvas(t *testing.T) {
	c := New(10, 20)

	if c.Width != 10 || c.Height != 20 {
		t.Error("values are wrong")
	}

	for i := range c.Pixels {
		for _, tup := range c.Pixels[i] {
			if !tpv.SameTuple(tup, tpv.Color(0, 0, 0)) {
				t.Fatal("not equal!")
			}
		}
	}
}

func TestWritePixel(t *testing.T) {
	c := New(10, 20)

	red := tpv.Color(1, 0, 0)

	if err := c.WritePixel(2, 3, red); err != nil {
		t.Error(err)
	}
	if err := c.WritePixel(10, 20, red); err == nil {
		t.Error("this error should not be nil")
	}
	if err := c.WritePixel(0, 0, red); err != nil {
		t.Error(err)
	}
	if err := c.WritePixel(-1, 0, red); err == nil {
		t.Error("this error should not be nil")
	}
	_, err := c.PixelAt(10, 20)
	if err == nil {
		t.Error("this error should not be nil")
	}
	_, err = c.PixelAt(-1, 0)
	if err == nil {
		t.Error("this error should not be nil")
	}

	tup, err := c.PixelAt(2, 3)
	if err != nil {
		t.Error(err)
	}

	if !tpv.SameTuple(tup, red) {
		t.Error("not the same color")
	}

	tup, err = c.PixelAt(0, 0)
	if err != nil {
		t.Error("this error should not be nil")
	}
	if !tpv.SameTuple(tup, red) {
		t.Error("not the same color")
	}
}

func TestPPMHeader(t *testing.T) {
	c := New(5, 3)

	c1 := tpv.Color(1.5, 0, 0)
	c2 := tpv.Color(0, 0.5, 0)
	c3 := tpv.Color(-0.5, 0, 1)

	if err := c.WritePixel(0, 0, c1); err != nil {
		t.Error(err)
	}
	if err := c.WritePixel(2, 1, c2); err != nil {
		t.Error(err)
	}
	if err := c.WritePixel(4, 2, c3); err != nil {
		t.Error(err)
	}

	ppm := c.ToPPM()

	expected := `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	if expected != ppm {
		t.Errorf("%q, should be equal to %q", expected, ppm)
	}

	c = New(10, 2)

	for rIndex, row := range c.Pixels {
		for cIndex := range row {
			c.Pixels[rIndex][cIndex] = tpv.Color(1, 0.8, 0.6)
		}
	}

	ppm = c.ToPPM()

	expected = `P3
10 2
255
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153
`

	if expected != ppm {
		t.Errorf("%q, should be equal to %q", expected, ppm)
	}
}
