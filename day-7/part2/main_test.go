package main

import (
	"sort"
	"testing"
)

func TestParse(t *testing.T) {
	a := parseLine(`32T3K 765`)
	if a.Bid != 765 {
		t.Errorf("expected bid to be 765 and got %d", a.Bid)
	}
	if len(a.Cards) != 5 {
		t.Errorf("expected 5 cards and got %d", len(a.Cards))
	}
	if a.Cards[0] != "3" && a.Cards[4] != "K" {
		t.Errorf("expected cards to be [3 2 T 3 K] and got %v", a.Cards)
	}
}

func TestSort(t *testing.T) {
	cards := []string{
		"K",
		"2",
		"A",
		"T",
		"8",
	}

	sort.Sort(ByCard(cards))
	if cards[0] != "A" && cards[4] != "2" {
		t.Errorf("expected cards to be [A K T 8 2] and got %v", cards)
	}
}

type HandTest struct {
	Cards         []string
	ExpectedValue int
}

func TestAnalyseWithoutJacks(t *testing.T) {
	tests := []HandTest{
		{
			Cards:         []string{"A", "A", "A", "A", "A"},
			ExpectedValue: 7,
		},
		{
			Cards:         []string{"A", "A", "A", "A", "Q"},
			ExpectedValue: 6,
		},
		{
			Cards:         []string{"A", "K", "K", "K", "K"},
			ExpectedValue: 6,
		},
		{
			Cards:         []string{"A", "A", "A", "K", "K"},
			ExpectedValue: 5,
		},
		{
			Cards:         []string{"A", "A", "K", "K", "K"},
			ExpectedValue: 5,
		},
		{
			Cards:         []string{"A", "A", "A", "K", "Q"},
			ExpectedValue: 4,
		},
		{
			Cards:         []string{"A", "K", "K", "K", "Q"},
			ExpectedValue: 4,
		},
		{
			Cards:         []string{"A", "K", "Q", "Q", "Q"},
			ExpectedValue: 4,
		},
		{
			Cards:         []string{"A", "A", "K", "K", "Q"},
			ExpectedValue: 3,
		},
		{
			Cards:         []string{"A", "A", "K", "Q", "Q"},
			ExpectedValue: 3,
		},
		{
			Cards:         []string{"A", "K", "K", "Q", "Q"},
			ExpectedValue: 3,
		},
		{
			Cards:         []string{"A", "A", "K", "Q", "T"},
			ExpectedValue: 2,
		},
		{
			Cards:         []string{"A", "K", "K", "Q", "T"},
			ExpectedValue: 2,
		},
		{
			Cards:         []string{"A", "K", "Q", "Q", "T"},
			ExpectedValue: 2,
		},
		{
			Cards:         []string{"A", "K", "Q", "T", "T"},
			ExpectedValue: 2,
		},
		{
			Cards:         []string{"A", "K", "Q", "T", "9"},
			ExpectedValue: 1,
		},
	}
	for i := 0; i < len(tests); i++ {
		v := analyseHand(tests[i].Cards)
		if v != tests[i].ExpectedValue {
			t.Errorf("expected value %d for hands %v but got value %d\n", tests[i].ExpectedValue, tests[i].Cards, v)
		}
	}
}

func TestAnalyseWithJacks(t *testing.T) {
	tests := []HandTest{
		{
			Cards:         []string{"T", "3", "T", "3", "J"},
			ExpectedValue: 5, // full house
		},
		{
			Cards:         []string{"2", "2", "3", "3", "J"},
			ExpectedValue: 5, // full house
		},
		{
			Cards:         []string{"T", "3", "J", "3", "J"},
			ExpectedValue: 6, // four of a kind
		},
		{
			Cards:         []string{"J", "3", "J", "3", "J"},
			ExpectedValue: 7, // five of a kind
		},
		{
			Cards:         []string{"2", "3", "4", "5", "A"},
			ExpectedValue: 1,
		},
		{
			Cards:         []string{"1", "2", "3", "4", "5"},
			ExpectedValue: 1,
		},
		{
			Cards:         []string{"1", "2", "3", "4", "4"},
			ExpectedValue: 2,
		},
		{
			Cards:         []string{"1", "2", "3", "4", "J"},
			ExpectedValue: 2,
		},
		{
			Cards:         []string{"1", "2", "2", "3", "3"},
			ExpectedValue: 3,
		},
		{
			Cards:         []string{"1", "2", "3", "3", "3"},
			ExpectedValue: 4,
		},
		{
			Cards:         []string{"1", "2", "3", "3", "J"},
			ExpectedValue: 4,
		},
		{
			Cards:         []string{"1", "2", "3", "J", "J"},
			ExpectedValue: 4,
		},
		{
			Cards:         []string{"1", "1", "2", "2", "2"},
			ExpectedValue: 5,
		},
		{
			Cards:         []string{"1", "1", "2", "2", "J"},
			ExpectedValue: 5,
		},
		{
			Cards:         []string{"1", "2", "2", "2", "2"},
			ExpectedValue: 6,
		},
		{
			Cards:         []string{"1", "2", "2", "2", "J"},
			ExpectedValue: 6,
		},
		{
			Cards:         []string{"1", "2", "2", "J", "J"},
			ExpectedValue: 6,
		},
		{
			Cards:         []string{"1", "2", "J", "J", "J"},
			ExpectedValue: 6,
		},
		{
			Cards:         []string{"1", "1", "1", "1", "1"},
			ExpectedValue: 7,
		},
		{
			Cards:         []string{"1", "1", "1", "1", "J"},
			ExpectedValue: 7,
		},
		{
			Cards:         []string{"1", "1", "1", "J", "J"},
			ExpectedValue: 7,
		},
		{
			Cards:         []string{"1", "1", "J", "J", "J"},
			ExpectedValue: 7,
		},
		{
			Cards:         []string{"1", "J", "J", "J", "J"},
			ExpectedValue: 7,
		},
		{
			Cards:         []string{"J", "J", "J", "J", "J"},
			ExpectedValue: 7,
		},
	}
	for i := 0; i < len(tests); i++ {
		v := analyseHand(tests[i].Cards)
		if v != tests[i].ExpectedValue {
			t.Errorf("expected value %d for hands %v but got value %d\n", tests[i].ExpectedValue, tests[i].Cards, v)
		}
	}
}

func TestRun(t *testing.T) {
	b := run(`KTJJT 34
2345A 1
Q2KJJ 13
Q2Q2Q 19
T3T3J 17
T3Q33 11
2345J 3
J345A 2
32T3K 5
T55J5 29
KK677 7
QQQJA 31
JJJJJ 37
JAAAA 43
AAAAJ 59
AAAAA 61
2AAAA 23
2JJJJ 53
JJJJ2 41`)
	if b != 6839 {
		t.Errorf("expected 6839 and got %d", b)
	}

	a := run(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)
	if a != 5905 {
		t.Errorf("expected 5905 and got %d", a)
	}
}
