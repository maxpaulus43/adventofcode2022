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
		cmp1, cmp2 := runeSet(line[:length/2]), runeSet(line[length/2:])
		intersection := cmp1.intersection(cmp2)
		sum += score(intersection[0])
	}
	return sum
}

func day3Part2() int {
	lines := linesFromFile("inputs/day3.txt")
	sum := 0
	for i := 0; i < len(lines); i += 3 {
		firstElf := runeSet(lines[i])
		secondElf := runeSet(lines[i+1])
		thirdElf := runeSet(lines[i+2])
		intersection := runeSet(string(firstElf.intersection(secondElf))).intersection(thirdElf)
		sum += score(intersection[0])
	}
	return sum
}
