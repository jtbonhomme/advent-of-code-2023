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

func TestIsStartingNode(t *testing.T) {
	var a bool
	a = isStartingNode(`AAA`)
	if !a {
		t.Errorf("expected AAA to be a starting node %v", a)
	}
	a = isStartingNode(`ABA`)
	if !a {
		t.Errorf("expected ABA to be a starting node %v", a)
	}
	a = isStartingNode(`CBA`)
	if !a {
		t.Errorf("expected CBA to be a starting node %v", a)
	}
	a = isStartingNode(`CBX`)
	if a {
		t.Errorf("expected CBX not to be a starting node %v", a)
	}
	a = isStartingNode(`CXX`)
	if a {
		t.Errorf("expected CXX not to be a starting node %v", a)
	}
	a = isStartingNode(`XXX`)
	if a {
		t.Errorf("expected XBX not to be a starting node %v", a)
	}
}

func TestIsEndingNodes(t *testing.T) {
	type test struct {
		nodes    []string
		expected bool
	}

	tests := []test{
		{
			nodes:    []string{},
			expected: false,
		},
		{
			nodes:    []string{"AAA"},
			expected: false,
		},
		{
			nodes:    []string{"AAA", "AAZ"},
			expected: false,
		},
		{
			nodes:    []string{"AAA", "AAA", "AAZ"},
			expected: false,
		},
		{
			nodes:    []string{"AAA", "AAZ", "AAZ"},
			expected: false,
		},
		{
			nodes:    []string{"AAZ", "AAZ", "AAZ"},
			expected: true,
		},
		{
			nodes:    []string{"AAZ", "AAZ"},
			expected: true,
		},
		{
			nodes:    []string{"AAZ"},
			expected: true,
		},
	}

	for _, tt := range tests {
		got := isEndingNodes(tt.nodes)
		if got != tt.expected {
			t.Errorf("expected %v ending node test to return %v, but got %v", tt.nodes, tt.expected, got)
		}
	}
}

func TestRun(t *testing.T) {
	a := run(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`)
	if a != 6 {
		t.Errorf("expected 6 and got %d", a)
	}

}
