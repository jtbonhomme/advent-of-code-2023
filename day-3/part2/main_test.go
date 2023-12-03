package main

import (
	"testing"
)

/*
In this schematic, there are two gears. The first is in the top left;
it has part numbers 467 and 35, so its gear ratio is 16345.
The second gear is in the lower right; its gear ratio is 451490.
(The * adjacent to 617 is not a gear because it is only adjacent to one part number.)
Adding up all of the gear ratios produces 467835.
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

func TestRun(t *testing.T) {
	a := run(i)
	if a != 467835 {
		t.Errorf("expected 467835 and got %d", a)
	}

}
