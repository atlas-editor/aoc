package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quasiPower(x, n int) int {
	if n == 0 {
		return 0
	}
	result := 1
	for i := 1; i < n; i++ {
		result *= x
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	cardNo := 1
	pileMap := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		matches := 0

		pileMap[cardNo]++

		_, nums, _ := strings.Cut(line, ":")
		firstPart, secondPart, _ := strings.Cut(nums, "|")

		winNums := strings.Fields(firstPart)
		winNumsSet := map[string]bool{}
		for _, n := range winNums {
			winNumsSet[n] = true
		}

		myNums := strings.Fields(secondPart)

		for _, n := range myNums {
			if winNumsSet[n] {
				matches++
			}
		}

		// part 2
		for j := 1; j <= matches; j++ {
			pileMap[cardNo+j] += pileMap[cardNo]
		}

		sum += quasiPower(2, matches)
		cardNo++
	}
	sumPart2 := 0
	for k, v := range pileMap {
		if k >= cardNo {
			continue
		}
		sumPart2 += v
	}
	fmt.Println(sum)
	fmt.Println(sumPart2)
}
