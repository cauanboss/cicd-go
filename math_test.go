package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	actual := Soma(15, 15)
	if actual != 30 {
		t.Errorf("Error actual = %v, and Expected = %v.", actual, 30)
	}
}
