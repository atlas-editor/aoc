package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func parsePropagation(info []string) []int {
	n0, _ := strconv.Atoi(info[0])
	n1, _ := strconv.Atoi(info[1])
	n2, _ := strconv.Atoi(info[2])
	return []int{n0, n1, n2}
}

func propagateNums(nums []int, propagationMap [][]int) []int {
	res := []int{}
OuterLoop:
	for _, num := range nums {
		for _, m := range propagationMap {
			dest := m[0]
			src := m[1]
			len := m[2]
			if src <= num && num < src+len {
				res = append(res, dest+num-src)
				continue OuterLoop
			}
		}
		res = append(res, num)
	}

	return res
}

func minFirstCoordinate(nums [][]int) int {
	m := nums[0][0]

	for _, n := range nums {
		if n[0] < m {
			m = n[0]
		}
	}

	return m
}

func splitRange(n0 int, n1 int, s0 int, s1 int) [][]int {
	c0 := min(max(n0, s0), n1)
	c1 := max(min(n1, s1), n0)
	splitPts := []int{n0, c0, c1, n1}

	splitPtsUniq := []int{}
	for _, pt := range splitPts {
		if !slices.Contains(splitPtsUniq, pt) {
			splitPtsUniq = append(splitPtsUniq, pt)
		}
	}

	res := [][]int{}
	for i := 0; i < len(splitPtsUniq)-1; i++ {
		res = append(res, []int{splitPtsUniq[i], splitPtsUniq[i+1]})
	}

	return res
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

func propagateRanges(nums [][]int, propagationMap [][]int) [][]int {
	allSplits := nums
	for _, m := range propagationMap {
		src := m[1]
		len := m[2]
		allSplits = splitAll(allSplits, src, src+len)
	}

	res := [][]int{}
OuterLoop:
	for _, n := range allSplits {
		n0 := n[0]
		n1 := n[1]
		for _, m := range propagationMap {
			dest := m[0]
			src := m[1]
			len := m[2]
			if src <= n0 && n1 <= src+len {
				n0Dest := dest + n0 - src
				n1Dest := n0Dest + n1 - n0
				res = append(res, []int{n0Dest, n1Dest})
				continue OuterLoop
			}
		}
		res = append(res, []int{n0, n1})
	}
	return res
}

func part1() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstLine := scanner.Text()
	_, seedStr, _ := strings.Cut(firstLine, ":")
	seedStrFields := strings.Fields(seedStr)
	seeds := []int{}

	for _, s := range seedStrFields {
		v, _ := strconv.Atoi(s)
		seeds = append(seeds, v)
	}

	propagationMap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		numFields := strings.Fields(line)
		if len(numFields) == 3 {
			propagationMap = append(propagationMap, parsePropagation(numFields))
		} else if len(propagationMap) > 0 {
			seeds = propagateNums(seeds, propagationMap)
			propagationMap = [][]int{}
		}
	}
	if len(propagationMap) > 0 {
		seeds = propagateNums(seeds, propagationMap)
	}
	fmt.Println(minInts(seeds))
}

func part2() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	firstLine := scanner.Text()
	_, seedStr, _ := strings.Cut(firstLine, ":")
	seedStrFields := strings.Fields(seedStr)
	seeds := [][]int{}

	for i := 0; i <= len(seedStrFields)/2; i += 2 {
		v, _ := strconv.Atoi(seedStrFields[i])
		r, _ := strconv.Atoi(seedStrFields[i+1])
		seeds = append(seeds, []int{v, v + r})
	}
	propagationMap := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		numFields := strings.Fields(line)
		if len(numFields) == 3 {
			propagationMap = append(propagationMap, parsePropagation(numFields))
		} else if len(propagationMap) > 0 {
			seeds = propagateRanges(seeds, propagationMap)
			propagationMap = [][]int{}
		}
	}
	if len(propagationMap) > 0 {
		seeds = propagateRanges(seeds, propagationMap)
	}
	fmt.Println(minFirstCoordinate(seeds))
}

func main() {
	part2()
}
