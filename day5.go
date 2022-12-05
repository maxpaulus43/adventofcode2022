package main

import (
	"strconv"
	"strings"
	"unicode"
)

func parseDrawing(pic []string) []stack[string] {
	result := make([]stack[string], 0)
	for j, ch := range pic[len(pic)-1] {
		if unicode.IsDigit(ch) {
			stack := make(stack[string], 0)
			for i := len(pic) - 2; i >= 0; i-- {
				if unicode.IsLetter(rune(pic[i][j])) {
					stack.push(string(pic[i][j]))
				}
			}
			result = append(result, stack)
		}
	}
	return result
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func day5Part1() string {
	fileStr := stringFromFile("inputs/day5.txt")
	sections := strings.Split(fileStr, "\n\n")
	stacks := parseDrawing(strings.Split(sections[0], "\n"))
	moves := strings.Split(sections[1], "\n")

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

// // almost exactly the same as part 1, but use a tempStack to reverse the order of the moved packages.
func day5Part2() string {
	fileStr := stringFromFile("inputs/day5.txt")
	sections := strings.Split(fileStr, "\n\n")
	stacks := parseDrawing(strings.Split(sections[0], "\n"))
	moves := strings.Split(sections[1], "\n")
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
