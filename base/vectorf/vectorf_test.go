package vectorf_test

import (
	"fmt"
	"strings"
	"testing"
	. "ubongo/base/vectorf"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	v := V{1.5, 2.3, 3.8}
	expected := "(1.5,2.3,3.8)"
	actual := strings.ReplaceAll(fmt.Sprint(v), "0", "")
	assert.Equal(t, expected, actual)
}

func TestMax(t *testing.T) {
	assert.Equal(t, 1.5, V{1.5, 0, -5}.Max())
	assert.Equal(t, -5.0, V{-10.5, -5.0, -50}.Max())
	assert.Equal(t, 99.0, V{1.5, 0, 99.0}.Max())
}

func TestAdd(t *testing.T) {
	a := V{1.5, 0, -5}
	b := V{-1, 2, 3.5}
	expected := V{0.5, 2, -1.5}
	assert.Equal(t, expected, a.Add(b))
}

func TestSub(t *testing.T) {
	a := V{1.5, 0, -5}
	b := V{-1, 2, 3.5}
	expected := V{2.5, -2, -8.5}
	assert.Equal(t, expected, a.Sub(b))
}

func TestMult(t *testing.T) {
	a := V{1.5, 0, -5}
	b := 2.0
	expected := V{3.0, 0, -10.0}
	assert.Equal(t, expected, a.Mult(b))
}

func TestDiv(t *testing.T) {
	a := V{1.5, 0, -5}
	b := 2.0
	expected := V{0.75, 0, -2.5}
	assert.Equal(t, expected, a.Div(b))
}

func TestFlip(t *testing.T) {
	a := V{1.5, 0, -5}
	expected := V{-1.5, 0, 5.0}
	assert.Equal(t, expected, a.Flip())
}
