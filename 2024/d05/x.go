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
	correct, _ := solve(input)
	return correct
}

func p2(input string) int {
	_, incorrect := solve(input)
	return incorrect
}

func solve(input string) (int, int) {
	parts := strings.Split(input, "\n\n")

	rules := set[pt]{}
	for _, r := range strings.Split(parts[0], "\n") {
		tmp := strings.Split(r, "|")
		rules[pt{atoi(tmp[0]), atoi(tmp[1])}] = true
	}

	correct := 0
	incorrect := 0
	for _, p := range strings.Split(parts[1], "\n") {
		nums := []int{}
		for _, tmp := range strings.Split(p, ",") {
			nums = append(nums, atoi(tmp))
		}

		sorted := sortByRules(nums, rules)
		middle := sorted[len(sorted)/2]
		if slices.Equal(nums, sorted) {
			correct += middle
		} else {
			incorrect += middle
		}
	}

	return correct, incorrect
}

func sortByRules(nums []int, rules set[pt]) []int {
	numsClone := slices.Clone(nums)
	slices.SortFunc(numsClone, func(a, b int) int {
		if v, ok := rules[pt{a, b}]; v && ok {
			return -1
		}
		return 1
	})

	return numsClone
}

/*
utils
*/

type pt [2]int

type set[T comparable] map[T]bool

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
