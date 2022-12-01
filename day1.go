package main

import (
	"sort"
	"strconv"
)

func day1Part1() int {
	input := linesFromFile("inputs/day1.txt")
	mostCalories := 0
	currElfCalories := 0

	for _, cal := range input {
		if len(cal) == 0 {
			currElfCalories = 0
		} else {
			c, err := strconv.Atoi(cal)
			check(err)
			currElfCalories += c
			if currElfCalories > mostCalories {
				mostCalories = currElfCalories
			}
		}
	}

	return mostCalories
}

func day1Part2() int {
	input := linesFromFile("inputs/day1.txt")
	caloriesPerElf := make([]int, 0)

	currElfCalories := 0
	for _, cal := range input {
		if len(cal) == 0 {
			caloriesPerElf = append(caloriesPerElf, currElfCalories)
			currElfCalories = 0
		} else {
			c, err := strconv.Atoi(cal)
			check(err)
			currElfCalories += c
		}
	}

	sort.Ints(caloriesPerElf)

	return sum(caloriesPerElf[len(caloriesPerElf)-3:]...)
}
