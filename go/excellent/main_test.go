package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	r := EvenOrOdd(10)
	if r != "even" {
		t.Errorf("expected: even, actual: %s", r)
	}
}
