package main

func day8Part1() int {
	lines := linesFromFile("inputs/day8.txt")
	grid := makeGridFromLines(lines)
	treesVisible := len(grid)*2 + len(grid[0])*2 - 4
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if isTreeVisible(i, j, grid) {
				treesVisible++
			}
		}
	}
	return treesVisible
}

func day8Part2() int {
	lines := linesFromFile("inputs/day8.txt")
	grid := makeGridFromLines(lines)
	maxScore := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			currScore := scoreTree(i, j, grid)
			if currScore > maxScore {
				maxScore = currScore
			}
		}
	}
	return maxScore
}

func makeGridFromLines(lines []string) [][]int {
	grid := make([][]int, 0, len(lines))
	for _, line := range lines {
		row := make([]int, 0, len(line))
		for _, ch := range line {
			n := int(ch - '0')
			row = append(row, n)
		}
		grid = append(grid, row)
	}
	return grid
}

func isTreeVisible(i, j int, grid [][]int) bool {
	// could possibly optimize with another array that keeps track of is visible
	// kinda dynamic programming approach
	currHeight := grid[i][j]
	// check north
	isVisibleFromNorth := true
	for k := i - 1; k >= 0; k-- {
		if grid[k][j] >= currHeight {
			isVisibleFromNorth = false
			break
		}
	}
	// check south
	isVisibleFromSouth := true
	for k := i + 1; k < len(grid); k++ {
		if grid[k][j] >= currHeight {
			isVisibleFromSouth = false
			break
		}
	}
	// check east
	isVisibleFromEast := true
	for k := j + 1; k < len(grid[0]); k++ {
		if grid[i][k] >= currHeight {
			isVisibleFromEast = false
			break
		}
	}
	// check west
	isVisibleFromWest := true
	for k := j - 1; k >= 0; k-- {
		if grid[i][k] >= currHeight {
			isVisibleFromWest = false
			break
		}
	}
	return isVisibleFromNorth ||
		isVisibleFromSouth ||
		isVisibleFromEast ||
		isVisibleFromWest
}

func scoreTree(i, j int, grid [][]int) int {
	currHeight := grid[i][j]
	// check north
	northScore := 0
	for k := i - 1; k >= 0; k-- {
		northScore++
		if grid[k][j] >= currHeight {
			break
		}
	}
	// check south
	southScore := 0
	for k := i + 1; k < len(grid); k++ {
		southScore++
		if grid[k][j] >= currHeight {
			break
		}
	}
	// check east
	eastScore := 0
	for k := j + 1; k < len(grid[0]); k++ {
		eastScore++
		if grid[i][k] >= currHeight {
			break
		}
	}
	// check west
	westScore := 0
	for k := j - 1; k >= 0; k-- {
		westScore++
		if grid[i][k] >= currHeight {
			break
		}
	}
	return northScore * southScore * eastScore * westScore
}
