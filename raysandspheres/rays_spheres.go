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
	Intersect(ray) collection
}

type sphere struct {
	center tpv.Tuple
	radio  float64
	id     string
}

type collection []intersection

// the object field will more than likely change in the future
type intersection struct {
	TValue float64
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

func NewIntersection(tValue float64, i intersectable) intersection {
	return intersection{tValue, i}
}

func NewSphere() *sphere {
	guid := xid.New()

	return &sphere{tpv.Point(0, 0, 0), 1, guid.String()}
}

func (s *sphere) Intersect(r ray) collection {
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

	return NewCollection(intersection{t1, s}, intersection{t2, s})
}

func NewCollection(inters ...intersection) collection {
	return append([]intersection{}, inters...)
}

func (c collection) Hit() (intersection, error) {
	hit := NewIntersection(math.MaxFloat64, nil)
	err := fmt.Errorf("no valid hit found in collection")

	for _, inter := range c {
		if inter.TValue >= 0 && inter.TValue < hit.TValue {
			hit = inter
			err = nil
		}
	}

	return hit, err
}
