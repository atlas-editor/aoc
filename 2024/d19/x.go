package main

import (
	"fmt"
	"os"
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
	c, _ := p(input)
	return c
}

func p2(input string) int {
	_, s := p(input)
	return s
}

func p(input string) (int, int) {
	parts := strings.Split(input, "\n\n")
	available := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1], "\n")

	c := 0
	s := 0
	for _, d := range designs {
		tmp := solve(available, d)
		if tmp > 0 {
			c++
		}
		s += tmp
	}

	return c, s
}

func solve(available []string, design string) int {
	var f func(string) int
	cache := map[string]int{}
	f = func(s string) int {
		if len(s) == 0 {
			return 1
		}

		if v, ok := cache[s]; ok {
			return v
		}

		possible := 0
		for _, a := range available {
			ss, ok := strings.CutPrefix(s, a)
			if ok {
				possible += f(ss)
			}
		}

		cache[s] = possible
		return possible
	}

	return f(design)
}
