package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]

	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(input string) int {
	lines := strings.Split(input, "\n")
	left, right := []int{}, []int{}

	for _, line := range lines {
		tmp := strings.Fields(line)
		left = append(left, atoi(tmp[0]))
		right = append(right, atoi(tmp[1]))
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0
	for i := range len(left) {
		distance += abs(left[i] - right[i])
	}

	return distance
}

func p2(input string) int {
	lines := strings.Split(input, "\n")
	left, right := []int{}, map[int]int{}

	for _, line := range lines {
		tmp := strings.Fields(line)
		left = append(left, atoi(tmp[0]))
		right[atoi(tmp[1])]++
	}

	similarity := 0
	for _, n := range left {
		similarity += n * right[n]
	}

	return similarity
}

/*
utils
*/

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
