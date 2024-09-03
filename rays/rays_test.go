package rays

import (
	"fmt"
	tpv "ray-tracer/tuplespointsvectors"
	"testing"
)

func TestNew(t *testing.T) {
	origin := tpv.Point(1, 2, 3)
	direction := tpv.Vector(4, 5, 6)

	r, err := New(origin, direction)
	if err != nil {
		t.Error(err)
	}

	if !tpv.SameTuple(origin, r.Origin) {
		t.Error("tuples should be equal")
	}
	if !tpv.SameTuple(direction, r.Direction) {
		t.Error("tuples should be equal")
	}
}

func TestPosition(t *testing.T) {
	r, err := New(tpv.Point(2, 3, 4), tpv.Vector(1, 0, 0))
	if err != nil {
		t.Error(err)
	}
	tests := []struct {
		time     float64
		expected tpv.Tuple
	}{{time: 0, expected: tpv.Point(2, 3, 4)},
		{time: 1, expected: tpv.Point(3, 3, 4)},
		{time: -1, expected: tpv.Point(1, 3, 4)},
		{time: 2.5, expected: tpv.Point(4.5, 3, 4)},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("time = %v", test.time), func(t *testing.T) {
			got := r.Position(test.time)
			if !tpv.SameTuple(got, test.expected) {
				t.Errorf("expected %v, but got %v", test.expected, got)
			}
		})
	}
}
