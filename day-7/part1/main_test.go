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

	sort.Sort(ByHand(cards))
	if cards[0] != "A" && cards[4] != "2" {
		t.Errorf("expected cards to be [A K T 8 2] and got %v", cards)
	}
}

func TestRun(t *testing.T) {
	a := run(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`)
	if a != 0 {
		t.Errorf("expected 0 and got %d", a)
	}
}
