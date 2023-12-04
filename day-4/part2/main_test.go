package main

import (
	"testing"
)

func TestSplitNumbers(t *testing.T) {
	n1 := splitNumbers("69 72 87 33 61 15  8 78 43 50")
	if len(n1) != 10 {
		t.Fatalf("expected n1 to be 10 number long, and got %d", len(n1))
	}
	if n1[0] != 69 {
		t.Errorf("expected n1[0] to be 69, and got %d", n1[0])
	}
	if n1[6] != 8 {
		t.Errorf("expected n1[6] to be 8, and got %d", n1[6])
	}
	if n1[9] != 50 {
		t.Errorf("expected n1[6] to be 8, and got %d", n1[9])
	}
}

func TestParse(t *testing.T) {
	line1 := "Card   1: 69 72 87 33 61 15  8 78 43 50 | 96 33 86 53 15 82 50 85 61  8 98 72 43 63 45 78 87 69 10 34 73 88 65 27 19"
	a, b := parseLine(line1)
	if len(a) != 10 {
		t.Errorf("expected winning numbers to be 10 long and got %d", len(a))
	}
	if len(b) != 25 {
		t.Errorf("expected my numbers to be 25 long and got %d", len(b))
	}
}

func TestRun(t *testing.T) {
	var lines = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	a := run(lines, 6)
	if a != 30 {
		t.Errorf("expected 30 and got %d", a)
	}
}
