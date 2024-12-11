package main

import (
	"fmt"
	"os"
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
	return solve(parse(input), 25)
}

func p2(input string) int {
	return solve(parse(input), 75)
}

func parse(input string) []int {
	nums := []int{}
	for _, t := range strings.Fields(input) {
		nums = append(nums, atoi(t))
	}
	return nums
}

func blink(n int) []int {
	switch {
	case n == 0:
		return []int{1}
	case len(strconv.Itoa(n))%2 == 0:
		nStr := strconv.Itoa(n)
		l := len(nStr) / 2
		return []int{atoi(nStr[:l]), atoi(nStr[l:])}
	default:
		return []int{n * 2024}
	}
}

func solve(nums []int, depth int) int {
	var dp func(int, int) int
	cache := map[[2]int]int{}

	dp = func(n, d int) int {
		if v, ok := cache[[2]int{n, d}]; ok {
			return v
		}

		b := blink(n)
		if d == 1 {
			return len(b)
		}

		s := 0
		for _, m := range b {
			s += dp(m, d-1)
		}

		cache[[2]int{n, d}] = s
		return s
	}

	r := 0
	for _, num := range nums {
		r += dp(num, depth)
	}

	return r
}

/*
utils
*/

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
