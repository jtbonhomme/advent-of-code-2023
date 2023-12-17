package main

import (
	"testing"
)

func TestDirectionParse(t *testing.T) {
	a := parseDirectionLine(`RRLLRLR`)
	if len(a) != 7 {
		t.Errorf("expected len to be 7 and got %d", len(a))
	}
	if a[0] != "R" && a[3] != "L" && a[4] != "R" && a[5] != "L" && a[6] != "R" {
		t.Errorf("expected directions to be [R R L L R L R] and got %v", a)
	}
}

func TestNavigationParse(t *testing.T) {
	a := parseNavigationLine(`AAA = (BBB, CCC)`)
	if a.Left != "BBB" {
		t.Errorf("expected left to be BBB and got %s", a.Left)
	}
	if a.Right != "CCC" {
		t.Errorf("expected right to be CCC and got %s", a.Right)
	}
}

func TestRun(t *testing.T) {
	a := run(`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`)
	if a != 2 {
		t.Errorf("expected 2 and got %d", a)
	}

	b := run(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`)
	if b != 6 {
		t.Errorf("expected 6 and got %d", b)
	}
}
