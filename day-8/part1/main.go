package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

type Node struct {
	Root  string
	Left  string
	Right string
}

func (n Node) Go(dir string) string {
	if dir == "R" {
		return n.Right
	}

	return n.Left
}

// AAA = (BBB, CCC) > Node{Root: "AAA", Left: "BBB", Right: "CCC"}
func parseNavigationLine(line string) Node {
	var n Node

	re := regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`) // want to know what is in front of 'at'
	res := re.FindAllStringSubmatch(line, -1)

	n.Root = res[0][1]
	n.Left = res[0][2]
	n.Right = res[0][3]

	return n
}

// RRRLLLR > [R R R L L L R]
func parseDirectionLine(line string) []string {
	var dir []string

	// parse directions: RRRLLLR
	for _, c := range []byte(line) {
		dir = append(dir, string(c))
	}

	return dir
}

func navigate(directionMap map[string]Node, next string, current int, directions []string) int {
	var step int

	if next == "ZZZ" {
		return current
	}

	for _, d := range directions {
		n := directionMap[next]
		next = n.Go(d)
		step++
	}

	return navigate(directionMap, next, current+step, directions)
}

func run(i string) int {
	var answer int
	var directionMap = make(map[string]Node)

	scanner := bufio.NewScanner(strings.NewReader(i))

	scanner.Scan()
	directions := parseDirectionLine(scanner.Text())

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		n := parseNavigationLine(line)
		directionMap[n.Root] = n
	}

	answer = navigate(directionMap, "AAA", answer, directions)

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
