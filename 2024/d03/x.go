package main

import (
	"fmt"
	"os"
	"regexp"
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
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	s := 0
	for _, m := range re.FindAllStringSubmatch(input, -1) {
		s += atoi(m[1]) * atoi(m[2])
	}

	return s
}

func p2(input string) int {
	re := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))`)

	enabled := true
	s := 0
	for _, m := range re.FindAllStringSubmatch(input, -1) {
		switch m[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				s += atoi(m[2]) * atoi(m[3])
			}
		}
	}

	return s
}

/*
utils
*/

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
