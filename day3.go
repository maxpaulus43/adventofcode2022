package main

func score(ch rune) int {
	switch {
	case int(ch) > int('a'):
		return int(ch) - int('a') + 1
	default:
		return int(ch) - int('A') + 27
	}
}

func day3Part1() int {
	lines := linesFromFile("inputs/day3.txt")
	sum := 0
	for _, line := range lines {
		length := len(line)
		cmp1, cmp2 := stringSet(line[:length/2]), stringSet(line[length/2:])
		union := cmp1.union(cmp2)
		sum += score(union[0])
	}
	return sum
}

func day3Part2() int {
	lines := linesFromFile("inputs/day3.txt")
	sum := 0
	for i := 0; i < len(lines); i += 3 {
		firstElf := stringSet(lines[i])
		secondElf := stringSet(lines[i+1])
		thirdElf := stringSet(lines[i+2])
		union := stringSet(string(firstElf.union(secondElf))).union(thirdElf)
		sum += score(union[0])
	}
	return sum
}
