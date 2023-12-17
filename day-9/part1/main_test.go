package main

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	a := parseLine(`0 3 6 9 12 15`)
	if len(a) != 6 {
		t.Errorf("expected len to be 6 and got %d", len(a))
	}
	if a[0] != 0 && a[1] != 3 && a[2] != 6 && a[4] != 12 && a[5] != 15 {
		t.Errorf("expected int to be [0 3 6 9 12 15] and got %v", a)
	}
}

func TestIsZero(t *testing.T) {
	a := isZero([]int{1, 2, 3, 0, 0})
	if a {
		t.Errorf("expected isZero to be false and got %v", a)
	}
	b := isZero([]int{0, 0, 0})
	if !b {
		t.Errorf("expected isZero to be true and got %v", b)
	}
}

func TestNext(t *testing.T) {
	a := next([]int{0, 3, 6, 9, 12, 15})
	if a != 18 {
		t.Errorf("expected next value to be 18 and got %d", a)
	}

	b := next([]int{19, 33, 49, 63, 72, 74, 68, 54, 33, 7, -21, -47, -66, -72, -58, -16, 63, 189, 373, 627, 964})
	if b != 1398 {
		t.Errorf("expected next value to be 1398 and got %d", b)
	}
}

func TestRun(t *testing.T) {
	a := run(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`)
	if a != 114 {
		t.Errorf("expected 114 and got %d", a)
	}
}
