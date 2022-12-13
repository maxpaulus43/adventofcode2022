package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day10Part1() int {
	lines := linesFromFile("inputs/day10.txt")
	q := make([]func(), 0)
	xValue := 1
	cycle := 1

	for _, line := range lines {
		fields := strings.Fields(line)
		switch fields[0] {
		case "noop":
			q = append(q, func() { cycle++ })
		case "addx":
			n, err := strconv.Atoi(fields[1])
			check(err)
			q = append(q, func() { cycle++ })
			q = append(q, func() {
				cycle++
				xValue += n
			})
		}
	}

	signalSum := 0
	for _, instr := range q {
		instr()
		switch cycle {
		case 20, 60, 100, 140, 180, 220:
			signalSum += cycle * xValue
		}
	}

	return signalSum
}

func day10Part2() int {
	lines := linesFromFile("inputs/day10.txt")
	q := make([]func(), 0)
	xValue := 1
	cycle := 1

	for _, line := range lines {
		fields := strings.Fields(line)
		switch fields[0] {
		case "noop":
			q = append(q, func() { cycle++ })
		case "addx":
			n, err := strconv.Atoi(fields[1])
			check(err)
			q = append(q, func() { cycle++ })
			q = append(q, func() {
				cycle++
				xValue += n
			})
		}
	}

	screen := strings.Builder{}
	for _, instr := range q {
		ch := ' '
		if abs(((cycle-1)%40)-xValue) < 2 {
			ch = '#'
		}
		screen.WriteRune(ch)
		if cycle%40 == 0 {
			screen.WriteRune('\n')
		}
		instr()
	}

	fmt.Print(screen.String())

	return 0
}
