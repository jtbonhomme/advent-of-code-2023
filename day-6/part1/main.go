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

func parseLine(line string) []int {
	var res = []int{}

	// line: 79 14 55 13
	lineString := strings.Split(line, ":")
	lineSplit := strings.Split(lineString[1], " ")
	// parse input
	for _, s := range lineSplit {
		v, err := strconv.Atoi(s)
		if err == nil {
			res = append(res, v)
		}
	}

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
	var races = []Race{}
	fmt.Println("parse input file")
	var answer int

	scanner := bufio.NewScanner(strings.NewReader(i))
	timeLine, ok := nextLine(scanner)
	if !ok {
		return -1
	}

	times := parseLine(timeLine)

	distanceLine, ok := nextLine(scanner)
	if !ok {
		return -1
	}

	distances := parseLine(distanceLine)

	for i, t := range times {
		races = append(races, Race{
			Time:     t,
			Distance: distances[i],
		})
	}

	fmt.Println("parse the races")
	for _, race := range races {
		var wins int
		for t := 0; t < race.Time; t++ {
			d := distance(race.Time, t)
			if d > race.Distance {
				wins++
			}
		}
		if answer == 0 {
			answer = wins
		} else {
			answer *= wins
		}
		fmt.Printf("race %d,%d have %d wins - cumulated answer is %d\n", race.Time, race.Distance, wins, answer)
	}

	return answer
}

func main() {
	var answer int
	answer = run(input)

	fmt.Println("Answer: ", answer)
}
