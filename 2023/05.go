package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func minInts(nums []int) int {
	m := nums[0]
	for _, n := range nums {
		if n < m {
			m = n
		}
	}
	return m
}

func minPairs(nums [][]int) int {
	m := nums[0][0]

	for _, n := range nums {
		if n[0] < m {
			m = n[0]
		}
	}

	return m
}

func splitRange(n0 int, n1 int, s0 int, s1 int) [][]int {
	if s0 <= n0 && s1 >= n1 {
		return [][]int{[]int{n0, n1}}
	}

	firstSplit := max(n0, s0)
	secondSplit := min(n1, s1)

	if firstSplit > n1 || secondSplit < n0 {
		return [][]int{[]int{n0, n1}}
	}

	if firstSplit == n0 {
		return [][]int{[]int{n0, secondSplit}, []int{secondSplit + 1, n1}}
	}

	if firstSplit > n0 {
		if secondSplit >= n1 {
			return [][]int{[]int{n0, firstSplit - 1}, []int{firstSplit, n1}}
		} else {
			return [][]int{[]int{n0, firstSplit - 1}, []int{firstSplit, secondSplit}, []int{secondSplit + 1, n1}}
		}
	}
	return [][]int{}
}

func splitAll(ranges [][]int, s0 int, s1 int) [][]int {
	splits := [][]int{}

	for _, n := range ranges {
		n0 := n[0]
		n1 := n[1]

		split := splitRange(n0, n1, s0, s1)
		splits = append(splits, split...)
	}

	return splits
}

func propagateInfo(nums [][]int, propagationMap [][]int) [][]int {
	if len(propagationMap) == 0 {
		return nums
	}
	allSplits := nums
	for _, m := range propagationMap {
		allSplits = splitAll(allSplits, m[1], m[1]+m[2]-1)
	}

	found := false
	result := [][]int{}
	// fmt.Println(allSplits)
	for _, n := range allSplits {
		found = false
		for _, m := range propagationMap {
			next := m[0]
			s0 := m[1]
			s1 := m[1] + m[2] - 1
			// fmt.Println("here", n, m)
			if s0 <= n[0] && s1 >= n[1] {
				n0 := next + n[0] - s0
				n1 := n0 + n[1] - n[0]
				result = append(result, []int{n0, n1})
				found = true
				// fmt.Println(m)
				// fmt.Println(result)
				break
			}
			// fmt.Println()
		}

		if !found {
			result = append(result, []int{n[0], n[1]})
		}
	}

	return result

	// for _, n := range nums {

	// 	// found = false
	// 	// for _, m := range propagationMap {
	// 	// 	if n >= m[1] && n < m[1]+m[2] {
	// 	// 		result = append(result, m[0]+n-m[1])
	// 	// 		found = true
	// 	// 	}
	// 	// }
	// 	// if !found {
	// 	// 	result = append(result, n)
	// 	// }
	// }

	// return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstLine := scanner.Text()
	_, seedStr, _ := strings.Cut(firstLine, ":")
	seedStrFields := strings.Fields(seedStr)
	seeds := [][]int{}

	for i := 0; i <= len(seedStrFields)/2; i += 2 {
		v, _ := strconv.Atoi(seedStrFields[i])
		r, _ := strconv.Atoi(seedStrFields[i+1])
		seeds = append(seeds, []int{v, v + r - 1})
	}
	// fmt.Println(seeds)
	propagationMap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()

		numFields := strings.Fields(line)
		if len(numFields) == 3 {
			n0, _ := strconv.Atoi(numFields[0])
			n1, _ := strconv.Atoi(numFields[1])
			n2, _ := strconv.Atoi(numFields[2])
			nums := []int{n0, n1, n2}
			propagationMap = append(propagationMap, nums)
		} else {
			seeds = propagateInfo(seeds, propagationMap)
			propagationMap = [][]int{}
		}
	}
	if len(propagationMap) > 0 {
		seeds = propagateInfo(seeds, propagationMap)
	}
	fmt.Println(minPairs(seeds))
}
