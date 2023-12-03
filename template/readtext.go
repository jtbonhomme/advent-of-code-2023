package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var answer int
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		answer++
		line := scanner.Text()
		fmt.Println(line)
	}
	fmt.Println("Answer: ", answer)
}
