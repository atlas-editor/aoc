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
	return solve(input, false)
}

func p2(input string) int {
	return solve(input, true)
}

type result struct {
	ok     bool
	target int
}

func solve(input string, useConcat bool) int {
	c := make(chan result)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		tmp := strings.Split(line, ": ")
		target := atoi(tmp[0])
		nums := []int{}
		for _, n := range strings.Fields(tmp[1]) {
			nums = append(nums, atoi(n))
		}
		go func() {
			c <- result{isPossible(target, nums, useConcat), target}
		}()
	}

	s := 0
	for range len(lines) {
		if r := <-c; r.ok {
			s += r.target
		}
	}
	return s
}

func isPossible(target int, nums []int, useConcat bool) bool {
	var f func(int, int) bool
	f = func(idx, total int) bool {
		if total > target {
			return false
		}
		if idx == len(nums) {
			return total == target
		}

		return f(idx+1, total+nums[idx]) || f(idx+1, total*nums[idx]) || (useConcat && f(idx+1, concat(total, nums[idx])))
	}

	return f(1, nums[0])
}

func concat(a, b int) int {
	return atoi(fmt.Sprintf("%v%v", a, b))
}

/*
utils
*/

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
