package vectorf_test

import (
	"fmt"
	"strings"
	"testing"
	. "ubongo/base/vectorf"

	"github.com/stretchr/testify/assert"
)

func TestVectorfString(t *testing.T) {
	v := V{1.5, 2.3, 3.8}
	expected := "(1.5,2.3,3.8)"
	actual := strings.ReplaceAll(fmt.Sprint(v), "0", "")
	assert.Equal(t, expected, actual)
}

func TestVectorfMax(t *testing.T) {
	a := V{1.5, 0, -5}
	m := a.Max()
	assert.Equal(t, 1.5, m)
}

func TestVectorfAdd(t *testing.T) {
	a := V{1.5, 0, -5}
	b := V{-1, 2, 3.5}
	expected := V{0.5, 2, -1.5}
	assert.Equal(t, expected, a.Add(b))
}

func TestVectorfSub(t *testing.T) {
	a := V{1.5, 0, -5}
	b := V{-1, 2, 3.5}
	expected := V{2.5, -2, -8.5}
	assert.Equal(t, expected, a.Sub(b))
}

func TestVectorfDiv(t *testing.T) {
	a := V{1.5, 0, -5}
	b := 2.0
	expected := V{0.75, 0, -2.5}
	assert.Equal(t, expected, a.Div(b))
}
func TestVectorfMult(t *testing.T) {
	a := V{1.5, 0, -5}
	b := 2.0
	expected := V{3.0, 0, -10.0}
	assert.Equal(t, expected, a.Mult(b))
}

func TestVectorfFlip(t *testing.T) {
	a := V{1.5, 0, -5}
	expected := V{-1.5, 0, 5.0}
	assert.Equal(t, expected, a.Flip())
}
