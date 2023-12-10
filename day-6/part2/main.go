package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseLine(line string) int {
	var res int
	var number string
	// line: 79 14 55 13
	lineString := strings.Split(line, ":")
	lineSplit := strings.Split(lineString[1], " ")
	// parse input
	for _, s := range lineSplit {
		if s != "" {
			number = number + s
		}
	}

	res, _ = strconv.Atoi(number)
	return res
}

type Race struct {
	Time     int
	Distance int
}

func distance(duration, boost int) int {
	return (duration - boost) * boost
}

func nextLine(scanner *bufio.Scanner) (string, bool) {
	ok := scanner.Scan()
	if !ok {
		return "", false
	}
	return scanner.Text(), true
}

func run(i string) int {
	var race = Race{}
	fmt.Println("parse input file")
	var answer int

	scanner := bufio.NewScanner(strings.NewReader(i))
	timeLine, ok := nextLine(scanner)
	if !ok {
		return -1
	}

	time := parseLine(timeLine)

	distanceLine, ok := nextLine(scanner)
	if !ok {
		return -1
	}

	dist := parseLine(distanceLine)

	race = Race{
		Time:     time,
		Distance: dist,
	}

	var minDuration int
	fmt.Println("parse the races")
	for t := 0; t < race.Time; t++ {
		d := distance(race.Time, t)
		if d > race.Distance {
			minDuration = t
			break
		}
	}

	fmt.Println("minDuration:", minDuration)
	answer = race.Time - minDuration*2 + 1

	return answer
}

func main() {
	var answer int
	answer = run(input)

	fmt.Println("Answer: ", answer)
}
