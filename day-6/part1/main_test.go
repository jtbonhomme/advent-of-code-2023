package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	a := parseLine(`line1: 1  23    1   6`)
	if len(a) != 4 {
		t.Errorf("expected 4 numbers and got %d", len(a))
	}

	if a[0] != 1 {
		t.Errorf("expected answer value at index 0 to be 1 and got %d", a[0])
	}

	if a[1] != 23 {
		t.Errorf("expected answer value at index 1 to be 23 and got %d", a[0])
	}

	if a[3] != 6 {
		t.Errorf("expected answer value at index 3 to be 6 and got %d", a[0])
	}
}

func TestRun(t *testing.T) {
	a := run(`Time:      7  15   30
Distance:  9  40  200`)
	if a != 288 {
		t.Errorf("expected 288 and got %d", a)
	}
}
