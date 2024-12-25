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
	nums := []int{}
	for _, line := range strings.Split(input, "\n") {
		nums = append(nums, atoi(line))
	}

	s := 0
	for _, n := range nums {
		for range 2000 {
			n = evolve(n)
		}
		s += n
	}

	return s
}

func p2(input string) int {
	nums := []int{}
	for _, line := range strings.Split(input, "\n") {
		nums = append(nums, atoi(line))
	}

	changesMap := map[[4]int][]int{}
	for _, n := range nums {
		last := []int{}
		last = append(last, n%10)

		changes := []int{}

		for i := range 2000 {
			n = evolve(n)
			last = append(last, n%10)
			changes = append(changes, last[i+1]-last[i])
		}

		seen := map[[4]int]bool{}
		for i := 3; i < len(changes); i++ {
			f := [4]int{changes[i-3], changes[i-2], changes[i-1], changes[i]}
			if _, ok := seen[f]; !ok {
				changesMap[f] = append(changesMap[f], last[i+1])
				seen[f] = true
			}
		}
	}

	maxVal := 0
	for _, v := range changesMap {
		s := 0
		for _, n := range v {
			s += n
		}
		if s > maxVal {
			maxVal = s
		}
	}

	return maxVal
}

func evolve(secret int) int {
	tmp := 64 * secret
	secret = secret ^ tmp
	secret = secret % 16777216

	tmp = secret / 32
	secret = secret ^ tmp
	secret = secret % 16777216

	tmp = 2048 * secret
	secret = secret ^ tmp
	secret = secret % 16777216

	return secret
}

/*
utils
*/

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
