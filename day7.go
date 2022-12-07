package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type file struct {
	parent *directory
	name   string
	size   int
}

type directory struct {
	name     string
	parent   *directory
	children []*directory
	files    []file
}

func newDir(name string, parent *directory) *directory {
	return &directory{
		name:     name,
		parent:   parent,
		children: make([]*directory, 0),
		files:    make([]file, 0),
	}
}

func (d *directory) getChild(childName string) (*directory, error) {
	for _, child := range d.children {
		if child.name == childName {
			return child, nil
		}
	}
	return d, errors.New("child doesn't exist")
}

func (d *directory) size() int {
	sum := 0
	for _, d := range d.children {
		sum += d.size()
	}
	for _, f := range d.files {
		sum += f.size
	}
	return sum
}

func (d *directory) addChildDir(childDir *directory) {
	d.children = append(d.children, childDir)
}
func (d *directory) addFile(name string, size int) {
	for _, file := range d.files {
		if file.name == name {
			return // if dir already has this file, return
		}
	}
	d.files = append(d.files, file{parent: d, name: name, size: size})
}

func buildDirectoryFromLines(lines []string) *directory {
	rootDir := newDir("/", nil)
	currDir := rootDir

	for _, line := range lines {
		fields := strings.Fields(line)

		switch fields[0] {
		case "$":
			cmd := fields[1]

			switch cmd {
			case "cd":
				toDir := fields[2]

				switch toDir {
				case "/":
					currDir = rootDir
				case "..":
					if currDir.name != rootDir.name {
						currDir = currDir.parent
					}
				default:
					childDir, err := currDir.getChild(toDir)

					if err != nil { // we haven't seen this dir yet, create it
						childDir = newDir(toDir, currDir)
						currDir.addChildDir(childDir)
						currDir = childDir
					} else {
						currDir = childDir
					}
				}
			}
		case "dir":
			dirName := fields[1]
			childDir, err := currDir.getChild(dirName)
			if err != nil { // we haven't seen this dir yet, create it
				childDir = newDir(dirName, currDir)
				currDir.addChildDir(childDir)
			}
		default: // this must be a file
			fileSize, err := strconv.Atoi(fields[0])
			check(err)
			currDir.addFile(fields[1], fileSize)
		}
	}
	return rootDir
}

func day7Part1() int {
	lines := linesFromFile("inputs/day7test.txt")
	rootDir := buildDirectoryFromLines(lines)
	printDirectory(rootDir, 0)
	return getSumOfTotalSizes(rootDir)
}

func day7Part2() int {
	lines := linesFromFile("inputs/day7.txt")
	rootDir := buildDirectoryFromLines(lines)
	totalSpace := 70000000
	freeSpaceNeeded := 30000000
	spaceUsed := rootDir.size()
	unusedSpace := totalSpace - spaceUsed
	spaceINeedToFree := freeSpaceNeeded - unusedSpace
	result := int(^uint(0) >> 1)

	return walk(spaceINeedToFree, result, rootDir)
}

func walk(spaceINeedToFree int, result int, dir *directory) int {
	size := dir.size()
	if size >= spaceINeedToFree && size < result {
		result = size
	}
	for _, child := range dir.children {
		result = walk(spaceINeedToFree, result, child)
	}
	return result
}

func getSumOfTotalSizes(dir *directory) int {
	sum := 0
	size := dir.size()
	if size <= 100000 {
		sum += size
	}
	for _, d := range dir.children {
		sum += getSumOfTotalSizes(d)
	}
	return sum
}

func printDirectory(d *directory, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	fmt.Printf("%sDir: %s, Size: %d\n", indent, d.name, d.size())
	for _, d := range d.children {
		printDirectory(d, level+1)
	}
	for _, f := range d.files {
		fmt.Printf("%s  File: %s, Size: %d\n", indent, f.name, f.size)
	}
}
