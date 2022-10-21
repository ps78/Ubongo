package vectorf

import (
	"fmt"
	"math"
)

type V [3]float64

func (a V) String() string {
	return fmt.Sprintf("(%f,%f,%f)", a[0], a[1], a[2])
}

func (a V) Max() float64 {
	return math.Max(math.Max(a[0], a[1]), a[2])
}

func (a V) Add(b V) V {
	return V{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func (a V) Sub(b V) V {
	return V{a[0] - b[0], a[1] - b[1], a[2] - b[2]}
}

func (a V) Div(b float64) V {
	return V{a[0] / b, a[1] / b, a[2] / b}
}

func (a V) Mult(b float64) V {
	return V{a[0] * b, a[1] * b, a[2] * b}
}

func (a V) Flip() V {
	return V{-a[0], -a[1], -a[2]}
}
