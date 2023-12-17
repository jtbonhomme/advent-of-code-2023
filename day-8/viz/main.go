package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
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

func isStartingNode(n string) bool {
	b := []byte(n)

	return string(b[2]) == "A"
}

func run(i string) int {
	var directionMap = make(map[string]Node)
	var startingNodes = []string{}

	f, err := os.Create("day-8/input.dot")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("digraph G {\n")
	if err != nil {
		panic(err)
	}

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
		var steps int
		node := startingNode
		for !isEndingNode(node) {
			for _, d := range directions {
				n := directionMap[node]
				steps++

				_, err = fmt.Fprintf(f, "  %s -> %s;\n", n.Root, n.Left)
				if err != nil {
					panic(err)
				}

				_, err = fmt.Fprintf(f, "  %s -> %s;\n", n.Root, n.Right)
				if err != nil {
					panic(err)
				}

				node = n.Go(d)
			}
		}

		_, err = fmt.Fprintf(f, "  %s [shape=Mdiamond];\n", startingNode)
		if err != nil {
			panic(err)
		}

		_, err = fmt.Fprintf(f, "  %s [shape=Msquare];\n", node)
		if err != nil {
			panic(err)
		}

		fmt.Printf("starting node %s resolves in %d steps\n", startingNode, steps)
	}

	_, err = f.WriteString("}\n")
	if err != nil {
		panic(err)
	}

	return 0
}

func main() {
	run(input)
}
