package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblemFactoryGet(t *testing.T) {
	f := GetProblemFactory()

	p := f.Get(Difficult, 12, 4)
	assert.NotNil(t, p)

	pnil := f.Get(Easy, 99, 1)
	assert.Nil(t, pnil)
}
