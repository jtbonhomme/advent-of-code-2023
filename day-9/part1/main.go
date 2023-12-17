package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func parseLine(line string) []int {
	var numbers = []int{}

	stringNumbers := strings.Split(line, " ")
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

func isZero(numbers []int) bool {
	for _, i := range numbers {
		if i != 0 {
			return false
		}
	}

	return true
}

func next(numbers []int) int {
	var last int
	var delta = []int{}

	fmt.Printf("%v\n", numbers)

	// end condition
	if isZero(numbers) {
		return 0
	}

	for i := 1; i < len(numbers); i++ {
		delta = append(delta, numbers[i]-numbers[i-1])
		last = numbers[i]
	}

	val := next(delta)

	return val + last
}

func run(i string) int {
	var answer int

	scanner := bufio.NewScanner(strings.NewReader(i))

	for scanner.Scan() {
		line := scanner.Text()
		arr := parseLine(line)
		nextNumber := next(arr)
		fmt.Println("next number is", nextNumber)
		answer += nextNumber
	}

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
