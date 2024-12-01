package main

import (
	"fmt"
	"os"
	"regexp"
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
		tmp := ints(line)
		left = append(left, tmp[0])
		right = append(right, tmp[1])
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
		tmp := ints(line)
		left = append(left, tmp[0])
		right[tmp[1]]++
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

func ints(s string) (r []int) {
	p := regexp.MustCompile(`-?\d+`)
	for _, e := range p.FindAllString(s, -1) {
		i, _ := strconv.Atoi(e)
		r = append(r, i)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
