package main

func day6Part1() int {
	line := stringFromFile("inputs/day6.txt")
	for i := 0; i < len(line)-4; i++ {
		set := runeSet(line[i : i+4])
		if len(set) == 4 {
			return i + 4
		}
	}
	return -1
}

func day6Part2() int {
	line := stringFromFile("inputs/day6.txt")
	for i := 0; i < len(line)-14; i++ {
		set := runeSet(line[i : i+14])
		if len(set) == 14 {
			return i + 14
		}
	}
	return -1
}
