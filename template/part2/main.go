package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func isDot(c byte) bool {
	return c == 46
}

func isDigit(c byte) bool {
	if c >= 48 && c <= 57 {
		return true
	}
	return false
}

func parseLine(line string, n int) int {
	var number, l int

	bline := []byte(line)
	fmt.Printf("%d: %s\n", n, line)

	// parse input
	for _, c := range bline {
		if isDigit(c) {
			number = number*10 + int(c-48)
			l++
			continue
		}
	}

	return 0
}

func run(i string) int {
	var answer, n int

	scanner := bufio.NewScanner(strings.NewReader(i))
	for scanner.Scan() {
		line := scanner.Text()
		res := parseLine(line, n)
		n++
		answer += res
		fmt.Println(line)
	}

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
