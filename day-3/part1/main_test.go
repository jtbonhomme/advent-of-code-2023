package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	line1 := "...780.685.....822....................560.35.............................529......780.....................453"
	n1, s1 := parseLine(line1, 0)
	if len(n1) != 8 {
		t.Errorf("expected 8 numbers and got %d", len(n1))
	}
	if len(s1) != 0 {
		t.Errorf("expected 0 symbols and got %d", len(s1))
	}
	line2 := "8...90*12..."
	n2, s2 := parseLine(line2, 0)
	if len(n2) != 3 {
		t.Errorf("expected 3 numbers and got %d", len(n2))
	}
	if len(s2) != 1 {
		t.Errorf("expected 1 symbols and got %d", len(s2))
	}
}

/*
In this schematic, two numbers are not part numbers because
they are not adjacent to a symbol: 114 (top right) and 58 (middle right).
Every other number is adjacent to a symbol and so is a part number;

	their sum is 4361.
*/
var i = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var j = `12.......*..
+.........34
.......-12..
..78........
..*....60...
78.........9
.5.....23..$
8...90*12...
............
2.2......12.
.*.........*
1.1..503+.56`

func TestRun(t *testing.T) {
	a := run(i)
	if a != 4361 {
		t.Errorf("expected 4361 and got %d", a)
	}

	b := run(j)
	if b != 925 {
		t.Errorf("expected 925 and got %d", b)
	}
}
