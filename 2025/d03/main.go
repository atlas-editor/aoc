package main

import (
	"fmt"
	"os"
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
		go func() {
			ch <- f([]byte(line), N)
		}()
	}

	r := 0
	for range len(lines) {
		r += <-ch
	}
	return r
}

func f(nums []byte, N int) int {
	a := 0
	b := len(nums) - (N - 1)

	s := 0
	for i := range N {
		m := byte(0)
		newa := 0
		for j, n := range nums {
			if j < a || j >= b {
				continue
			}
			if n > m {
				m = n
				newa = j + 1
			}
		}
		s += int(m-48) * pow(10, N-(i+1))

		a = newa
		b += 1
	}

	return s
}

/*
 * utils
 */

func pow(base, exp int) int {
	if exp < 0 {
		panic("exp must be non-negative")
	}
	r := 1
	for exp > 0 {
		if exp&1 == 1 {
			r *= base
		}
		base *= base
		exp >>= 1
	}
	return r
}
