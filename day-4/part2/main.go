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

func parseLine(line string) ([]int, []int) {
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

func score(winningNumbers, myNumbers []int) int {
	var score int
	for _, n := range myNumbers {
		if slices.Contains(winningNumbers, n) {
			score++
		}
	}

	return score
}

func displayCards(scratchcards []int) {
	for i := 0; i < len(scratchcards); i++ {
		fmt.Printf("Card %d: %d - ", i+1, scratchcards[i])
	}
	fmt.Println("")
}

func run(i string, n int) int {
	var answer, l int
	scratchcards := make([]int, n)
	// init number of scratchcards (eg 1)
	for i := 0; i < n; i++ {
		scratchcards[i] = 1
	}

	fmt.Println(" -- init --")
	displayCards(scratchcards)

	scanner := bufio.NewScanner(strings.NewReader(i))
	for scanner.Scan() {
		fmt.Printf("\n-- parse card %d -- \n", l+1)
		line := scanner.Text()
		fmt.Println(line)
		winningNumbers, myNumbers := parseLine(line)
		numberOfScratchcards := score(winningNumbers, myNumbers)
		fmt.Printf("score of card %d = %d \n", l+1, numberOfScratchcards)
		l++
		for i := l; i < l+numberOfScratchcards; i++ {
			scratchcards[i] += scratchcards[l-1]
		}
		displayCards(scratchcards)
	}

	for i := 0; i < len(scratchcards); i++ {
		answer += scratchcards[i]
	}

	return answer
}

func main() {
	answer := run(input, 196)
	fmt.Println("Answer: ", answer)
}
