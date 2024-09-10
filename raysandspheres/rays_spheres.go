package raysandspheres

import (
	"fmt"
	"math"

	tpv "ray-tracer/tuplespointsvectors"

	"github.com/rs/xid"
)

type ray struct {
	Origin, Direction tpv.Tuple
}

type intersectable interface {
	Intersect(ray) []intersection
}

type sphere struct {
	center tpv.Tuple
	radio  float64
	id     string
}

// the object field will more than likely change in the future
type intersection struct {
	T      float64
	Object intersectable
}

func NewRay(origin, direction tpv.Tuple) (ray, error) {
	if !origin.IsPoint() {
		return ray{}, fmt.Errorf("origin must be a point")
	}
	if !direction.IsVector() {
		return ray{}, fmt.Errorf("direction must be a vector")
	}
	nr := ray{origin, direction}

	return nr, nil
}

func (r ray) Position(t float64) tpv.Tuple {
	return tpv.Add(r.Origin, tpv.ScMult(r.Direction, t))
}

func NewIntersection(t float64, i intersectable) *intersection {
	return &intersection{t, i}
}

func NewSphere() *sphere {
	guid := xid.New()

	return &sphere{tpv.Point(0, 0, 0), 1, guid.String()}
}

func (s *sphere) Intersect(r ray) []intersection {
	sphereToRay := tpv.Subtract(r.Origin, s.center)

	a := tpv.Dot(r.Direction, r.Direction)
	b := 2 * tpv.Dot(r.Direction, sphereToRay)
	c := tpv.Dot(sphereToRay, sphereToRay) - 1

	discriminant := math.Pow(b, 2) - (4 * a * c)

	if discriminant < 0 {
		return []intersection{}
	}

	t1 := (-b - (math.Sqrt(discriminant))) / (2 * a)
	t2 := (-b + (math.Sqrt(discriminant))) / (2 * a)

	return Intersections(intersection{t1, s}, intersection{t2, s})
}

func Intersections(inters ...intersection) []intersection {
	return append([]intersection{}, inters...)
}
