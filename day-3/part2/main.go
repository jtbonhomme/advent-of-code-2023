package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

/*
		x	0123456789
	y

	0		467..114..
	1		...*......
	2		..35..633.
	3		......#...
	4		617*......
	5		.....+.58.
	6		..592.....
	7		......755.
	8		...$.*....
	0		.664.598..

	467 is at pos [0, 0] (end len 3)
	* is at pos [3, 1] (and len 1)
*/

type Position struct {
	X int
	Y int
}

type Number struct {
	Value int
	Len   int
	Pos   Position
}

type Gear struct {
	Pos Position
}

func isGear(c byte) bool {
	return c == 42
}

func isDot(c byte) bool {
	return c == 46
}

func isDigit(c byte) bool {
	if c >= 48 && c <= 57 {
		return true
	}
	return false
}

func parseLine(line string, n int) ([]Number, []Gear) {
	var gears = []Gear{}
	var numbers = []Number{}
	var number, l int

	bline := []byte(line)
	fmt.Printf("%d: %s\n", n, line)

	// parse input
	for i, c := range bline {
		if isDigit(c) {
			number = number*10 + int(c-48)
			l++
			continue
		}

		if isGear(c) {
			// the only other option is symbol
			gears = append(gears, Gear{
				Pos: Position{
					X: i,
					Y: n,
				},
			})
			fmt.Printf("%s [%d,%d] - ", string(c), i, n)
		}

		if number != 0 {
			numbers = append(numbers, Number{
				Value: number,
				Len:   l,
				Pos: Position{
					X: i - l,
					Y: n,
				},
			})
			fmt.Printf("%d (%d) [%d,%d] - ", number, l, i-l, n)
			l = 0
			number = 0
		}

	}

	// if last number in line
	if number != 0 {
		numbers = append(numbers, Number{
			Value: number,
			Len:   l,
			Pos: Position{
				X: len(line) - l,
				Y: n,
			},
		})
		fmt.Printf("%d (%d) [%d,%d] - ", number, l, len(line)-l, n)
	}

	fmt.Println("")

	return numbers, gears
}

func distance(p1, p2 Position) int {
	dist := int(math.Max(math.Abs(float64(p1.X-p2.X)), math.Abs(float64(p1.Y-p2.Y))))
	return dist
}

func isAdjacent(posSymbol, posNumber Position, l int) bool {
	for i := 0; i < l; i++ {
		posDigit := Position{X: posNumber.X + i, Y: posNumber.Y}
		if distance(posSymbol, posDigit) == 1 {
			return true
		}
	}
	return false
}

func run(i string) int {
	var gears = []Gear{}
	var numbers = []Number{}
	var answer, n int

	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		partnumbers, partgears := parseLine(line, n)
		n++
		numbers = append(numbers, partnumbers...)
		gears = append(gears, partgears...)
	}

	for _, g := range gears {
		ok, n1, n2 := hasAdjacents(g.Pos, numbers)
		if ok {
			answer += n1.Value * n2.Value
			fmt.Printf("gear * [%d, %d] has adjacents numbers %d and %d, accumulator is %d\n", g.Pos.X, g.Pos.Y, n1.Value, n2.Value, answer)
		}
	}

	return answer
}

func hasAdjacents(p Position, numbers []Number) (bool, Number, Number) {
	var n1, n2 Number
	for _, n := range numbers {
		if isAdjacent(p, n.Pos, n.Len) {
			if n1.Len == 0 {
				n1 = n
			} else {
				n2 = n
				return true, n1, n2
			}
		}
	}

	return false, n1, n2
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
