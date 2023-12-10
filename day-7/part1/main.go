package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"sort"
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
	v := cardsRank[c]
	return v
}

// ByCard implements sort.Interface for []string based on
// the Card facing value.
type ByCard []string

func (a ByCard) Len() int      { return len(a) }
func (a ByCard) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCard) Less(i, j int) bool {
	iValue := cardRank(a[i])
	jValue := cardRank(a[j])
	return iValue > jValue
}

func handScore(cards []string) int {
	var score int

	for i, c := range cards {
		v := cardRank(c)
		score += (len(cards) - i) * (len(cards) - i) * v * v
	}

	return score
}

// ByHandValue implements sort.Interface for []Hand based on
// the Value field.
type ByHandValue []Hand

func (a ByHandValue) Len() int      { return len(a) }
func (a ByHandValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByHandValue) Less(i, j int) bool {
	iValue := a[i].Value
	jValue := a[j].Value

	if iValue == jValue {
		for k, _ := range a[i].Cards {
			iScore := cardRank(a[i].Cards[k])
			jScore := cardRank(a[j].Cards[k])
			if iScore == jScore {
				continue
			}
			return iScore > jScore

		}
	}

	return iValue > jValue
}

// 7: AAAAA
// 6: AAAAQ or AKKKK
// 5: AAAKK or AAKKK
// 4: AAAKQ or AKKKQ or AKQQQ
// 3: AAKKQ or AAKQQ or AKKQQ
// 2: AAKQT or AKKQT or AKQQT or AKQTT
// 1: AKQT9
func analyseHand(cards []string) int {
	var hasThreeOfAKind bool
	var numberOfPairs int
	copycards := make([]string, len(cards))
	copy(copycards, cards)

	sort.Sort(ByCard(copycards))
	for i := 0; i < len(copycards); i++ {
		n := 1
		for j := i + 1; j < len(copycards); j++ {
			if copycards[i] == copycards[j] {
				n++
			}
		}
		if n == 5 {
			return 7
		}
		if n == 4 {
			return 6
		}
		if n == 3 {
			hasThreeOfAKind = true
		}
		// warning, three of a kind also counts as a pair
		if n == 2 {
			numberOfPairs++
		}
	}

	if hasThreeOfAKind && numberOfPairs == 2 {
		return 5
	}
	if hasThreeOfAKind && numberOfPairs == 1 {
		return 4
	}
	if !hasThreeOfAKind && numberOfPairs == 2 {
		return 3
	}
	if !hasThreeOfAKind && numberOfPairs == 1 {
		return 2
	}

	return 1
}

func run(i string) int {
	var answer int
	var hands = []Hand{}

	scanner := bufio.NewScanner(strings.NewReader(i))
	for scanner.Scan() {
		hand := parseLine(scanner.Text())
		hand.Value = analyseHand(hand.Cards)
		hands = append(hands, hand)
	}

	fmt.Println("sorted hands by value: ")
	sort.Sort(ByHandValue(hands))
	for i, h := range hands {
		rank := len(hands) - i
		answer += rank * h.Bid
		fmt.Printf("rank %d - %v (value %d - bid %d) - cumulated answer %d\n", rank, h.Cards, h.Value, h.Bid, answer)
	}

	return answer
}

func main() {
	answer := run(input)
	fmt.Println("Answer: ", answer)
}
