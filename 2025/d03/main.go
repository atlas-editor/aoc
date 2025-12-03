package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(input string) int {
	return p(input, 2)
}

func p2(input string) int {
	return p(input, 12)
}

func p(input string, N int) int {
	lines := strings.Split(input, "\n")
	ch := make(chan int, len(lines))
	for _, line := range lines {
		nums := []int{}
		for _, c := range line {
			nums = append(nums, int(c-48))
		}

		go func() {
			ch <- f(nums, N)
		}()
	}

	r := 0
	for range len(lines) {
		r += <-ch
	}
	return r
}

func f(nums []int, N int) int {
	a := 0
	b := len(nums) - (N - 1)

	s := ""
	for range N {
		m := -1
		newa := -1
		for j, n := range nums {
			if j < a || j >= b {
				continue
			}
			if n > m {
				m = n
				newa = j + 1
			}
		}
		s += strconv.Itoa(m)

		a = newa
		b += 1
	}

	return atoi(s)
}

/*
 * utils
 */

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
