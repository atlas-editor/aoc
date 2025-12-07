package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"utils"
)

func main() {
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(input string) int {
	res, _ := p(input)
	return res
}

func p2(input string) int {
	_, res := p(input)
	return res
}

func p(input string) (int, int) {
	area := 0
	ribbon := 0
	for _, line := range strings.Split(input, "\n") {
		nums := utils.Ints(line)
		slices.Sort(nums)
		a, b, c := nums[0], nums[1], nums[2]

		area += 2*a*b + 2*b*c + 2*c*a + min(a*b, b*c, c*a)
		ribbon += 2*a + 2*b + a*b*c
	}
	return area, ribbon
}
