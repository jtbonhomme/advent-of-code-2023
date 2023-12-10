package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var input string

type SeedRange struct {
	SeedRangeStart  int
	SeedRangeLength int
}

func parseSeeds(line string) []SeedRange {
	var res = []int{}
	var seedRanges []SeedRange

	// seeds: 79 14 55 13
	seedsString := strings.Split(line, ":")
	seedsSplit := strings.Split(seedsString[1], " ")
	// parse input
	for _, s := range seedsSplit {
		v, err := strconv.Atoi(s)
		if err == nil {
			res = append(res, v)
		}
	}

	for i := 0; i < len(res); i += 2 {
		seedRanges = append(seedRanges, SeedRange{
			SeedRangeStart:  res[i],
			SeedRangeLength: res[i+1],
		})
	}

	return seedRanges
}

type RangeMap struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

// x-to-y map:
// 50 98 2
// 52 50 48
func parseMap(i string) []*RangeMap {
	var m = []*RangeMap{}

	scanner := bufio.NewScanner(strings.NewReader(i))

	// header
	ok := scanner.Scan()
	if !ok {
		return m
	}
	// map line 50 98 2
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, " ")
		destinationRangeStart, _ := strconv.Atoi(ranges[0])
		sourceRangeStart, _ := strconv.Atoi(ranges[1])
		rangeLength, _ := strconv.Atoi(ranges[2])
		rangeMap := RangeMap{
			DestinationRangeStart: destinationRangeStart,
			SourceRangeStart:      sourceRangeStart,
			RangeLength:           rangeLength,
		}

		m = append(m, &rangeMap)
	}

	return m
}

func nextLine(scanner *bufio.Scanner) (string, bool) {
	ok := scanner.Scan()
	if !ok {
		return "", false
	}
	return scanner.Text(), true
}

func getMappedValue(m []*RangeMap, v int) int {
	for _, r := range m {
		if v >= r.SourceRangeStart && v < r.SourceRangeStart+r.RangeLength {
			return r.DestinationRangeStart + v - r.SourceRangeStart
		}
	}
	return v
}

func run(i string) int {
	var maps = [][]*RangeMap{}
	fmt.Println("parse input file")
	var answer int

	scanner := bufio.NewScanner(strings.NewReader(i))
	// seed line
	// seeds: 79 14 55 13
	seedLine, ok := nextLine(scanner)
	if !ok {
		return -1
	}

	seeds := parseSeeds(seedLine)

	// blank line
	_, ok = nextLine(scanner)
	if !ok {
		return -1
	}

	for ok {
		ll := ""
		line := ""
		line, ok = nextLine(scanner)
		for line != "" && ok {
			ll += line + "\n"
			line, ok = nextLine(scanner)
		}
		m := parseMap(ll)
		maps = append(maps, m)
	}

	fmt.Println("search lowest location")

	var minLocation int = math.MaxInt
	for _, seedRange := range seeds {
		for seed := seedRange.SeedRangeStart; seed < seedRange.SeedRangeStart+seedRange.SeedRangeLength; seed++ {
			soil := getMappedValue(maps[0], seed)
			fertilizer := getMappedValue(maps[1], soil)
			water := getMappedValue(maps[2], fertilizer)
			light := getMappedValue(maps[3], water)
			temperature := getMappedValue(maps[4], light)
			humidity := getMappedValue(maps[5], temperature)
			location := getMappedValue(maps[6], humidity)
			if location < minLocation {
				minLocation = location
				answer = location
			}
		}
	}

	return answer
}

func main() {
	start := time.Now()
	answer := run(input)

	fmt.Printf("Answer: %d (in %s)\n", answer, time.Since(start))
}
