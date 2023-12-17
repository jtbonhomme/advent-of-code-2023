package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"time"
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

	re := regexp.MustCompile(`([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)`) // want to know what is in front of 'at'
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

func isEndingNode(n string) bool {
	b := []byte(n)

	return string(b[2]) == "Z"
}

func isEndingNodes(nodes []string) bool {
	var res = true

	if len(nodes) == 0 {
		return false
	}

	for _, n := range nodes {
		res = res && isEndingNode(n)
	}

	return res
}

func navigate(directionMap map[string]Node, node string, directions []string) int {
	var step int

	for !isEndingNode(node) {
		for _, d := range directions {
			n := directionMap[node]
			node = n.Go(d)
			step++
		}
	}

	return step
}

func isStartingNode(n string) bool {
	b := []byte(n)

	return string(b[2]) == "A"
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(integers ...int) int {
	if len(integers) < 2 {
		return 0
	}
	a := integers[0]
	b := integers[1]
	result := a * b / GCD(a, b)
	integers = integers[2:]

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func run(i string) int {
	var answers []int
	var directionMap = make(map[string]Node)
	var startingNodes = []string{}

	scanner := bufio.NewScanner(strings.NewReader(i))

	scanner.Scan()
	directions := parseDirectionLine(scanner.Text())
	fmt.Printf("direction (%d): %v\n", len(directions), directions)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		n := parseNavigationLine(line)
		directionMap[n.Root] = n
		if isStartingNode(n.Root) {
			startingNodes = append(startingNodes, n.Root)
		}
	}

	for _, startingNode := range startingNodes {
		steps := navigate(directionMap, startingNode, directions)
		answers = append(answers, steps)
		fmt.Printf("starting node %s resolves in %d steps\n", startingNode, steps)
	}

	return LCM(answers...)
}

func main() {
	start := time.Now()
	answer := run(input)
	fmt.Printf("Answer: %d (in %v)\n", answer, time.Since(start))
}
