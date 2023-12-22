package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type brick struct {
	start [3]int
	end   [3]int
}

func brickCmp(b0, b1 brick) int {
	if b0.start[2] != b1.start[2] {
		return b0.start[2] - b1.start[2]
	}
	return b0.end[2] - b1.end[2]
}

func isBelow(b0, b1 brick) bool {
	if b0.end[2] >= b1.start[2] {
		return false
	}

	b0x0 := b0.start[0]
	b0x1 := b0.end[0]
	b0y0 := b0.start[1]
	b0y1 := b0.end[1]

	b1x0 := b1.start[0]
	b1x1 := b1.end[0]
	b1y0 := b1.start[1]
	b1y1 := b1.end[1]

	if b1x1 < b0x0 || b1x0 > b0x1 || b1y1 < b0y0 || b1y0 > b0y1 {
		return false
	}

	return true
}

func isDirectlyBelow(b0, b1 brick) bool {
	return isBelow(b0, b1) && b0.end[2]+1 == b1.start[2]
}

func isSubset(a, b []int) bool {
	for _, e := range a {
		if !slices.Contains(b, e) {
			return false
		}
	}
	return true
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	bricks := []brick{}
	for scn.Scan() {
		line := scn.Text()
		b := strings.Split(line, "~")

		brickStart := [3]int{}
		for idx, v := range strings.Split(b[0], ",") {
			i, _ := strconv.Atoi(v)
			brickStart[idx] = i
		}

		brickEnd := [3]int{}
		for idx, v := range strings.Split(b[1], ",") {
			i, _ := strconv.Atoi(v)
			brickEnd[idx] = i
		}

		br := brick{brickStart, brickEnd}
		bricks = append(bricks, br)
	}

	slices.SortFunc(bricks, brickCmp)

	stackedBricks := []brick{}
	for i := 0; i < len(bricks); i++ {
		z := 1
		for j := i - 1; j >= 0; j-- {
			if isBelow(stackedBricks[j], bricks[i]) {
				z = stackedBricks[j].end[2] + 1
				break
			}
		}
		fallenBrick := bricks[i]
		d := fallenBrick.start[2] - z
		fallenBrick.start[2] = z
		fallenBrick.end[2] = fallenBrick.end[2] - d

		stackedBricks = append(stackedBricks, fallenBrick)
		slices.SortFunc(stackedBricks, func(a, b brick) int { return a.end[2] - b.end[2] })
	}

	slices.SortFunc(stackedBricks, brickCmp)

	supBy := map[int][]int{}
	for i := 0; i < len(stackedBricks); i++ {
		for j := i - 1; j >= 0; j-- {
			if isDirectlyBelow(stackedBricks[j], stackedBricks[i]) {
				supBy[i] = append(supBy[i], j)
			}
		}
	}

	part1 := len(stackedBricks)
	part2 := 0
	for i := 0; i < len(stackedBricks); i++ {
		curr := []int{i}
		for j := i + 1; j < len(stackedBricks); j++ {
			if len(supBy[j]) > 0 && isSubset(supBy[j], curr) {
				curr = append(curr, j)
			}
		}
		part2 += len(curr) - 1
		if len(curr) > 1 {
			part1--
		}

	}

	fmt.Println(part1)
	fmt.Println(part2)
}
