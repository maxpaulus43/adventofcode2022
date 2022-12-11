package main

import (
	"fmt"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func (p pos) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func reconcile(head pos, tail *pos) string {
	dxdy := fmt.Sprintf("%d,%d", head.x-tail.x, head.y-tail.y)
	switch dxdy {
	case "-2,-1", "-1,-2", "-2,-2":
		tail.x--
		tail.y--
	case "0,-2":
		tail.y--
	case "1,-2", "2,-1", "2,-2":
		tail.x++
		tail.y--
	case "2,0":
		tail.x++
	case "2,1", "1,2", "2,2":
		tail.x++
		tail.y++
	case "0,2":
		tail.y++
	case "-1,2", "-2,1", "-2,2":
		tail.x--
		tail.y++
	case "-2,0":
		tail.x--
	}
	return tail.String()
}

func day9Part1() int {
	lines := linesFromFile("inputs/day9.txt")
	head := pos{0, 0}
	tail := pos{0, 0}
	visited := make(set[string], 0)
	visited.add(tail.String())

	for _, line := range lines {
		dir := line[0]
		distance, err := strconv.Atoi(strings.Fields(line)[1])
		check(err)

		for i := 0; i < distance; i++ {
			switch dir {
			case 'R':
				head.x++
			case 'U':
				head.y--
			case 'D':
				head.y++
			case 'L':
				head.x--
			}
			visited.add(reconcile(head, &tail))
		}
	}
	return len(visited)
}

func makeRope(n int) []pos {
	rope := make([]pos, 0, n)
	for i := 0; i < n; i++ {
		rope = append(rope, pos{0, 0})
	}
	return rope
}

func day9Part2() int {
	lines := linesFromFile("inputs/day9.txt")
	rope := makeRope(10)
	head := &rope[0]
	tail := rope[len(rope)-1]
	visited := make(set[string], 0)
	visited.add(tail.String())

	for _, line := range lines {
		dir := line[0]
		distance, err := strconv.Atoi(strings.Fields(line)[1])
		check(err)

		for i := 0; i < distance; i++ {
			switch dir {
			case 'R':
				head.x++
			case 'U':
				head.y--
			case 'D':
				head.y++
			case 'L':
				head.x--
			}
			lastVisited := ""
			for knot := 0; knot < len(rope)-1; knot++ {
				lastVisited = reconcile(rope[knot], &rope[knot+1])
			}
			visited.add(lastVisited)
		}
	}

	return len(visited)

}
