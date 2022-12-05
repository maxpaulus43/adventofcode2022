package main

import (
	"strconv"
	"strings"
)

func parseDrawing(lines []string) []stack[string] {
	result := make([]stack[string], 0)
	pic := make([][]rune, 0)

	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		pic = append(pic, []rune(line))
	}

	for j, ch := range pic[len(pic)-1] {
		if ch == ' ' || ch == '\n' {
			continue
		}
		stack := make(stack[string], 0)
		for i := len(pic) - 2; i >= 0; i-- {
			if pic[i][j] == ' ' {
				continue
			}
			stack.push(string(pic[i][j]))
		}
		result = append(result, stack)
	}
	return result
}

func parseMoves(lines []string) []string {
	picDone := false
	result := make([]string, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			picDone = true
			continue
		}
		if !picDone {
			continue
		}
		result = append(result, line)
	}
	return result
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func day5Part1() string {
	lines := linesFromFile("inputs/day5.txt")
	stacks := parseDrawing(lines)
	moves := parseMoves(lines)
	for _, move := range moves {
		if len(move) == 0 {
			continue
		}
		fields := strings.Fields(move)
		cnt := toInt(fields[1])
		from := toInt(fields[3]) - 1
		to := toInt(fields[5]) - 1

		for i := 0; i < cnt; i++ {
			stacks[to].push(stacks[from].pop())
		}
	}

	result := ""
	for _, s := range stacks {
		result += s.peek()
	}

	return result
}

// almost exactly the same as part 1, but use a tempStack to reverse the order of the moved packages.
func day5Part2() string {
	lines := linesFromFile("inputs/day5.txt")
	stacks := parseDrawing(lines)
	moves := parseMoves(lines)
	for _, move := range moves {
		if len(move) == 0 {
			continue
		}
		fields := strings.Fields(move)
		cnt := toInt(fields[1])
		from := toInt(fields[3]) - 1
		to := toInt(fields[5]) - 1
		// use this to reverse the order
		tempStack := make(stack[string], cnt)

		for i := 0; i < cnt; i++ {
			tempStack.push(stacks[from].pop())
		}
		for i := 0; i < cnt; i++ {
			stacks[to].push(tempStack.pop())
		}
	}

	result := ""
	for _, s := range stacks {
		result += s.peek()
	}

	return result
}
