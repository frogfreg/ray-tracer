package rays

import (
	"math"

	tpv "ray-tracer/tuplespointsvectors"

	"github.com/rs/xid"
)

type ray struct {
	Origin, Direction tpv.Tuple
}

type sphere struct {
	center tpv.Tuple
	radio  float64
	id     string
}

func NewRay(origin, direction tpv.Tuple) ray {

	nr := ray{origin, direction}

	return nr
}

func (r ray) Position(t float64) tpv.Tuple {
	return tpv.Add(r.Origin, tpv.ScMult(r.Direction, t))
}

func NewSphere() sphere {

	guid := xid.New()

	return sphere{tpv.NewPoint(0, 0, 0), 1, guid.String()}
}

func Intersect(s sphere, r ray) []float64 {
	sphereToRay := tpv.Subtract(r.Origin, s.center)

	a := tpv.Dot(r.Direction, r.Direction)
	b := 2 * tpv.Dot(r.Direction, sphereToRay)
	c := tpv.Dot(sphereToRay, sphereToRay) - 1

	discriminant := math.Pow(b, 2) - (4 * a * c)

	if discriminant < 0 {
		return []float64{}
	}

	t1 := (-b - (math.Sqrt(discriminant))) / (2 * a)
	t2 := (-b + (math.Sqrt(discriminant))) / (2 * a)

	return []float64{t1, t2}
}
