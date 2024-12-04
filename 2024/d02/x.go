package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	path := "input.txt"
	//path = "sample.txt"

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

		// 58  54  52  49  46  43
		//   -4  -2  -3  -3  -3

		if nums[0] == 58 {
			fmt.Println("")
		}

		b := false
		if isAlmostSafeIncreasing(nums) || isAlmostSafeDecreasing(nums) {
			safe++
			b = true
			//continue outer
		}

		for i := range len(nums) {
			report := slices.Delete(slices.Clone(nums), i, i+1)
			if isSafeIncreasing(report) || isSafeDecreasing(report) {
				safe++
				if !b {
					fmt.Println("not b")
					fmt.Println(nums)
					panic("")
				}
				continue outer
			}
		}

		if b {
			fmt.Println("b")
			fmt.Println(nums)
			panic("")
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

func isAlmostSafeIncreasing(nums []int) bool {
	var invalidIncr []int
	var diffs []int

	for i := range len(nums) - 1 {
		diff := nums[i+1] - nums[i]
		diffs = append(diffs, diff)

		if !(diff >= 1 && diff <= 3) {
			invalidIncr = append(invalidIncr, i)
		}
	}

	switch len(invalidIncr) {
	case 0:
		return true
	case 1:
		idx := invalidIncr[0]

		idxM1 := 0
		if idx != 0 {
			idxM1 = diffs[idx-1]
		}

		idxP0 := diffs[idx]

		idxP1 := 0
		if idx != len(diffs)-1 {
			idxP1 = diffs[idx+1]
		}
		if (idxM1+idxP0 >= 1 && idxM1+idxP0 <= 3) || (idxP0+idxP1 >= 1 && idxP0+idxP1 <= 3) {
			fmt.Println("one wrong", nums)
			return true
		}
	case 2:
		idx0, idx1 := invalidIncr[0], invalidIncr[1]
		if idx0 == idx1-1 && diffs[idx0]+diffs[idx0+1] >= 1 && diffs[idx0]+diffs[idx0+1] <= 3 {
			fmt.Println("two wrong", nums)
			return true
		}
	}

	return false
}

func isAlmostSafeDecreasing(nums []int) bool {
	slices.Reverse(nums)
	return isAlmostSafeIncreasing(nums)
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
