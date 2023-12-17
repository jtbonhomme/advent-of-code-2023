package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	MaxRed   int = 12
	MaxGreen int = 13
	MaxBlue  int = 14
)

func main() {
	var gameSplit = regexp.MustCompile(`^Game [0-9]+: |(?:(?P<number>[0-9]+) (?P<color>[a-z]+)[,;]?[ ]?)+$`)
	var colorSplit = regexp.MustCompile(`(?P<number>[0-9]+) (?P<color>[a-z]+)[,;]?[ ]?`)
	var answer, n int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var maxRed, maxGreen, maxBlue int
		n++
		line := scanner.Text()
		fmt.Println(line)
		res := gameSplit.FindAllStringSubmatch(line, -1)
		if len(res) > 1 {
			if len(res[1]) > 0 {
				res2 := colorSplit.FindAllStringSubmatch(res[1][0], -1)
				for _, c := range res2 {
					if len(c) == 3 {
						i, err := strconv.Atoi(c[1])
						if err != nil {
							panic(err)
						}
						switch c[2] {
						case "red":
							if i > maxRed {
								maxRed = i
							}
						case "green":
							if i > maxGreen {
								maxGreen = i
							}
						case "blue":
							if i > maxBlue {
								maxBlue = i
							}
						default:
						}
					}
				}
				fmt.Printf("\tmaxRed: %d, maxGreen: %d, maxBlue: %d\n", maxRed, maxGreen, maxBlue)
				if maxBlue > MaxBlue || maxGreen > MaxGreen || maxRed > MaxRed {
					continue
				}
				answer += n
			}
		}
	}
	fmt.Println("Answer: ", answer)
}
