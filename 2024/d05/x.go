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
	_, correct, _ := parse(input)

	s := 0
	for _, c := range correct {
		s += c[len(c)/2]
	}

	return s
}

func p2(input string) int {
	rules, _, incorrect := parse(input)

	for i := range incorrect {
		slices.SortFunc(incorrect[i], func(a, b int) int {
			for _, r := range rules {
				if r[0] == a && r[1] == b {
					return -1
				}
			}
			return 1
		})
	}

	s := 0
	for _, i := range incorrect {
		s += i[len(i)/2]
	}

	return s
}

func parse(input string) ([]pt, [][]int, [][]int) {
	parts := strings.Split(input, "\n\n")

	rules := []pt{}
	for _, r := range strings.Split(parts[0], "\n") {
		tmp := strings.Split(r, "|")
		rules = append(rules, pt{atoi(tmp[0]), atoi(tmp[1])})
	}

	correct := [][]int{}
	incorrect := [][]int{}
outer:
	for _, p := range strings.Split(parts[1], "\n") {
		nums := []int{}
		for _, tmp := range strings.Split(p, ",") {
			nums = append(nums, atoi(tmp))
		}
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				for _, r := range rules {
					if r[0] == nums[j] && r[1] == nums[i] {
						incorrect = append(incorrect, nums)
						continue outer
					}
				}
			}
		}
		correct = append(correct, nums)
	}

	return rules, correct, incorrect
}

/*
utils
*/

type pt [2]int

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
