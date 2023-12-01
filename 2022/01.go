package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// assume arr is sorted and non-empty
func insertNum(arr []int, num int) []int {
	if num > arr[0] {
		arr = append(arr[1:], num)
		sort.Ints(arr)
	}
	return arr
}

func sum(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	N := 3
	maxElves := make([]int, N)
	currElf := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			maxElves = insertNum(maxElves, currElf)
			currElf = 0
		} else {
			currCal, _ := strconv.Atoi(line)
			currElf += currCal
		}
	}
	maxElves = insertNum(maxElves, currElf)

	fmt.Println(sum(maxElves))
}
