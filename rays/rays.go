package rays

import (
	"fmt"
	tpv "ray-tracer/tuplespointsvectors"

	"github.com/rs/xid"
)

type ray struct {
	Origin, Direction tpv.Tuple
}

func NewRay(origin, direction tpv.Tuple) ray {

	nr := ray{origin, direction}

	return nr

}

func (r ray) Position(t float64) tpv.Tuple {
	return tpv.Add(r.Origin, tpv.ScMult(r.Direction, t))
}

func Sphere() string {
	guid := xid.New()

	fmt.Printf("%T\n", guid)

	return guid.String()

}
