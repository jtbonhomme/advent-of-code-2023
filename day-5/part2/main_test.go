package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	a := parseLine(`line1`, 1)
	if a != 0 {
		t.Errorf("expected 0 and got %d", a)
	}
}

func TestRun(t *testing.T) {
	a := run(`line1
line2`)
	if a != 0 {
		t.Errorf("expected 0 and got %d", a)
	}
}
