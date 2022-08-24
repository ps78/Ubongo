package main

import "testing"

func TestUbongo(t *testing.T) {
	result := 42
	expected := 42
	if result != expected {
		t.Errorf("result: %v\n expected: %v", result, expected)
	}
}
