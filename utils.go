package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func linesFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

func stringFromFile(fileName string) string {
	bytes, err := os.ReadFile(fileName)
	check(err)
	return string(bytes)
}

func intsFromFile(fileName string) []int {
	return stringsToInts(linesFromFile(fileName))
}

func stringsToInts(strings []string) []int {
	result := make([]int, 0, len(strings))
	for _, s := range strings {
		n, err := strconv.Atoi(s)
		check(err)
		result = append(result, n)
	}

	return result
}

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func abs[T number](n T) T {
	return T(math.Abs(float64(n)))
}

func sum[T number](nums ...T) T {
	var sum T
	sum = 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func avg[T number](nums ...T) float64 {
	return float64(sum(nums...)) / float64(len(nums))
}

type numbers[T number] []T

func (r numbers[T]) Len() int           { return len(r) }
func (r numbers[T]) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r numbers[T]) Less(i, j int) bool { return r[i] < r[j] }

type stack[T comparable] []T

func (s *stack[T]) pop() T {
	tmp := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return tmp
}
func (s *stack[T]) push(elem T) {
	*s = append(*s, elem)
}
func (s stack[T]) peek() T {
	return s[len(s)-1]
}

type set[T comparable] map[T]struct{}

func (s set[T]) has(value T) bool {
	_, ok := s[value]
	return ok
}
func (s set[T]) add(value T) {
	s[value] = struct{}{}
}
func (s set[T]) remove(value T) {
	delete(s, value)
}
func (s set[T]) intersection(otherSet set[T]) []T {
	result := make([]T, 0)
	for v := range s {
		if otherSet.has(v) {
			result = append(result, v)
		}
	}
	return result
}
func runeSet(str string) set[rune] {
	result := make(set[rune])
	for _, ch := range str {
		result.add(ch)
	}
	return result
}
