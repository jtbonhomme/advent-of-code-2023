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

func TestAnalyse(t *testing.T) {
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

func TestRun(t *testing.T) {
	a := run(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)
	if a != 6440 {
		t.Errorf("expected 6440 and got %d", a)
	}
}
