package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	a := parseLine(`line1: 1  23    1   6`)

	if a != 12316 {
		t.Errorf("expected answer value to be 12316 and got %d", a)
	}
}

func TestRun(t *testing.T) {
	a := run(`Time:      7  15   30
Distance:  9  40  200`)
	if a != 71503 {
		t.Errorf("expected 71503 and got %d", a)
	}
}
