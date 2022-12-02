package main

import (
	"strings"
)

var translate = map[string]string{"X": "A", "Y": "B", "C": "Z"}
var beats = map[string]string{"A": "B", "B": "C", "C": "A"}
var loses = map[string]string{"A": "C", "B": "A", "C": "B"}
var pointValueFor = map[string]int{"A": 1, "B": 2, "C": 3}

func day2Part1() int {
	lines := linesFromFile("inputs/day2.txt")
	totalScore := 0
	for _, line := range lines {
		roundScore := 0
		fields := strings.Fields(line)
		oppMove, myMove := fields[0], translate[fields[1]]

		roundScore += pointValueFor[myMove]

		if myMove == beats[oppMove] {
			roundScore += 6
		} else if myMove == oppMove {
			roundScore += 3
		}
		totalScore += roundScore
	}
	return totalScore
}

func day2Part2() int {
	lines := linesFromFile("inputs/day2.txt")
	totalScore := 0

	for _, line := range lines {
		roundScore := 0
		fields := strings.Fields(line)
		oppMove, endState := fields[0], fields[1]

		switch endState {
		case "X": // lose
			roundScore += pointValueFor[loses[oppMove]]
		case "Y": // draw
			roundScore += pointValueFor[oppMove]
			roundScore += 3
		case "Z": // win
			roundScore += pointValueFor[beats[oppMove]]
			roundScore += 6
		}
		totalScore += roundScore
	}
	return totalScore
}
