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

type Hand struct {
	Cards []string
	Bid   int
	Value int
}

var cardsRank map[string]int

func init() {
	cardsRank = make(map[string]int)
	cardsRank["A"] = 13
	cardsRank["K"] = 12
	cardsRank["Q"] = 11
	cardsRank["J"] = 10
	cardsRank["T"] = 9
	cardsRank["9"] = 8
	cardsRank["8"] = 7
	cardsRank["7"] = 6
	cardsRank["6"] = 5
	cardsRank["5"] = 4
	cardsRank["4"] = 3
	cardsRank["3"] = 2
	cardsRank["2"] = 1
}

func parseLine(line string) Hand {
	var hand Hand

	lineSplit := strings.Split(line, " ")

	// parse cards
	for _, c := range []byte(lineSplit[0]) {
		hand.Cards = append(hand.Cards, string(c))
	}

	hand.Bid, _ = strconv.Atoi(lineSplit[1])

	return hand
}

func cardRank(c string) int {
	v, _ := cardsRank[c]
	return v
}

// ByHand implements sort.Interface for []Person based on
// the Age field.
type ByHand []string

func (a ByHand) Len() int      { return len(a) }
func (a ByHand) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByHand) Less(i, j int) bool {
	iValue := cardRank(a[i])
	jValue := cardRank(a[j])
	return iValue > jValue
}

func isFiveOfAKind(hand Hand) bool {
	return false
}

func isFourOfAKind(hand Hand) bool {
	return false
}

func isFullHouse(hand Hand) bool {
	return false
}

func isThreeOfAKind(hand Hand) bool {
	return false
}

func isTwoPair(hand Hand) bool {
	return false
}

func isOnePair(hand Hand) bool {
	return false
}

func isHighCard(hand Hand) bool {
	return false
}

func value(hand Hand) int {
	return 0
}

func run(i string) int {
	var answer int
	var hands = []Hand{}

	scanner := bufio.NewScanner(strings.NewReader(i))
	for scanner.Scan() {
		hands = append(hands, parseLine(scanner.Text()))
	}

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
