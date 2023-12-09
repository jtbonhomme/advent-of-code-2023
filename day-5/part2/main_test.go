package main

import (
	"testing"
)

func TestParseSeeds(t *testing.T) {
	res := []SeedRange{
		{
			SeedRangeStart:  79,
			SeedRangeLength: 14,
		},
		{
			SeedRangeStart:  55,
			SeedRangeLength: 13,
		},
	}

	a := parseSeeds(`seeds: 79 14 55 13`)
	if len(a) != 2 {
		t.Errorf("expected int array len to be 4 and got %d", len(a))
	}
	for i, v := range a {
		if v != res[i] {
			t.Errorf("expected answer value at index %d to be %v and got %v", i, res[i], v)
		}
	}
}

func TestParseMap(t *testing.T) {
	input := `seed-to-soil map:
50 98 2
52 50 48`

	m := parseMap(input)
	if len(m) != 2 {
		t.Errorf("expected map len to be 2 and got %d", len(m))
	}

	if m[0].SourceRangeStart != 98 {
		t.Errorf("expected SourceRangeStart of map at index 0 to be 98 and got %d", m[0].DestinationRangeStart)
	}

	if m[0].DestinationRangeStart != 50 {
		t.Errorf("expected DestinationRangeStart of map at index 0 to be 50 and got %d", m[0].DestinationRangeStart)
	}

	if m[0].RangeLength != 2 {
		t.Errorf("expected RangeLength of map at index 0 to be 2 and got %d", m[0].DestinationRangeStart)
	}

	if m[1].SourceRangeStart != 50 {
		t.Errorf("expected SourceRangeStart of map at index 1 to be 50 and got %d", m[0].DestinationRangeStart)
	}

	if m[1].DestinationRangeStart != 52 {
		t.Errorf("expected DestinationRangeStart of map at index 1 to be 52 and got %d", m[0].DestinationRangeStart)
	}

	if m[1].RangeLength != 48 {
		t.Errorf("expected RangeLength of map at index 1 to be 48 and got %d", m[0].DestinationRangeStart)
	}
}

func TestGetMappedValue(t *testing.T) {
	// seed-to-soil map:
	// 50 98 2
	// 52 50 48
	m := []*RangeMap{
		{
			DestinationRangeStart: 50,
			SourceRangeStart:      98,
			RangeLength:           2,
		},
		{
			DestinationRangeStart: 52,
			SourceRangeStart:      50,
			RangeLength:           48,
		},
	}

	for i := 0; i < 50; i++ {
		v := getMappedValue(m, i)
		if i != v {
			t.Errorf("expected value %d to be mapped with %d, but got %d", i, i, v)
		}
	}

	for i := 50; i < 98; i++ {
		v := getMappedValue(m, i)
		if i+2 != v {
			t.Errorf("expected value %d to be mapped with %d, but got %d", i, i+2, v)
		}
	}

	if getMappedValue(m, 98) != 50 {
		t.Errorf("expected value 98 to be mapped with 50, but got %d", getMappedValue(m, 98))
	}

	if getMappedValue(m, 99) != 51 {
		t.Errorf("expected value 99 to be mapped with 51, but got %d", getMappedValue(m, 98))
	}

	for i := 100; i < 1000; i++ {
		v := getMappedValue(m, i)
		if i != v {
			t.Errorf("expected value %d to be mapped with %d, but got %d", i, i, v)
		}
	}

}

func TestRun(t *testing.T) {
	a := run(`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`)
	if a != 46 {
		t.Errorf("expected 46 and got %d", a)
	}
}
