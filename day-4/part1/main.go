package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"slices"
	"strconv"
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

func splitNumbers(s string) []int {
	var numbers = []int{}

	stringNumbers := strings.Split(s, " ")
	for _, sn := range stringNumbers {
		if sn == "" {
			continue
		}
		n, err := strconv.Atoi(sn)
		if err != nil {
			log.Printf("error converting %s: %s", sn, err.Error())
			continue
		}
		numbers = append(numbers, n)
	}

	return numbers
}

func parseLine(line string, n int) ([]int, []int) {
	payload := strings.Split(line, ":")
	if len(payload) != 2 {
		panic("line split failed")
	}
	payloadSplit := strings.Split(payload[1], "|")
	if len(payloadSplit) != 2 {
		panic("payload split failed")
	}

	winningNumbersStrings := payloadSplit[0]
	myNumbersStrings := payloadSplit[1]

	winningNumbers := splitNumbers(winningNumbersStrings)
	myNumbers := splitNumbers(myNumbersStrings)

	return winningNumbers, myNumbers
}

func updateScore(currentScore int) int {
	if currentScore == 0 {
		return 1
	}
	return currentScore * 2
}

func score(winningNumbers, myNumbers []int) int {
	var score int
	for _, n := range myNumbers {
		if slices.Contains(winningNumbers, n) {
			score = updateScore(score)
		}
	}
	return score
}

func run(i string) int {
	var answer, n int

	scanner := bufio.NewScanner(strings.NewReader(i))
	for scanner.Scan() {
		line := scanner.Text()
		winningNumbers, myNumbers := parseLine(line, n)
		answer += score(winningNumbers, myNumbers)
		n++
		fmt.Println(line)
	}

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
