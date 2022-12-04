package main

import "strings"

func contains(aRange []int, bRange []int) bool {
	return aRange[0] <= bRange[0] && aRange[1] >= bRange[1]
}

func overlap(aRange []int, bRange []int) bool {
	return aRange[1] >= bRange[0] && aRange[0] <= bRange[1]
}

func rangesFromLine(line string) ([]int, []int) {
	ranges := strings.Split(line, ",")
	aRange := stringsToInts(strings.Split(ranges[0], "-"))
	bRange := stringsToInts(strings.Split(ranges[1], "-"))
	return aRange, bRange
}

func day4Part1() int {
	lines := linesFromFile("inputs/day4.txt")
	pairCount := 0
	for _, line := range lines {
		aRange, bRange := rangesFromLine(line)
		if contains(aRange, bRange) || contains(bRange, aRange) {
			pairCount++
		}
	}
	return pairCount
}

func day4Part2() int {
	lines := linesFromFile("inputs/day4.txt")
	pairCount := 0
	for _, line := range lines {
		aRange, bRange := rangesFromLine(line)
		if overlap(aRange, bRange) {
			pairCount++
		}
	}
	return pairCount
}
