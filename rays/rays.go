package rays

import tpv "ray-tracer/tuplespointsvectors"

type ray struct {
	Origin, Direction tpv.Tuple
}

func NewRay(origin, direction tpv.Tuple) ray {

	nr := ray{origin, direction}

	return nr

}
