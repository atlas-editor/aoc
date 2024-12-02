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

	safe := 0
	for _, line := range lines {
		report := ints(line)
		if isSafeIncreasing(report) || isSafeDecreasing(report) {
			safe++
		}
	}

	return safe
}

func p2(input string) int {
	lines := strings.Split(input, "\n")

	safe := 0
outer:
	for _, line := range lines {
		nums := ints(line)
		for i := range len(nums) {
			report := slices.Delete(slices.Clone(nums), i, i+1)
			if isSafeIncreasing(report) || isSafeDecreasing(report) {
				safe++
				continue outer
			}
		}
	}

	return safe
}

func isSafeIncreasing(nums []int) bool {
	for i := range len(nums) - 1 {
		if !(nums[i+1]-nums[i] > 0 && nums[i+1]-nums[i] <= 3) {
			return false
		}
	}
	return true
}

func isSafeDecreasing(nums []int) bool {
	slices.Reverse(nums)
	return isSafeIncreasing(nums)
}

/*
utils
*/

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func ints(s string) (r []int) {
	tmp := strings.Fields(s)
	for _, t := range tmp {
		r = append(r, atoi(t))
	}
	return r
}
