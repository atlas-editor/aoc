//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type rock [][2]int

func buildRocks() []rock {
	r0 := rock{[2]int{0, 0}, [2]int{0, 1}, [2]int{0, 2}, [2]int{0, 3}}
	r1 := rock{[2]int{0, 0}, [2]int{1, -1}, [2]int{1, 0}, [2]int{1, 1}, [2]int{2, 0}}
	r2 := rock{[2]int{0, 0}, [2]int{0, 1}, [2]int{0, 2}, [2]int{1, 2}, [2]int{2, 2}}
	r3 := rock{[2]int{0, 0}, [2]int{1, 0}, [2]int{2, 0}, [2]int{3, 0}}
	r4 := rock{[2]int{0, 0}, [2]int{0, 1}, [2]int{1, 0}, [2]int{1, 1}}

	return []rock{r0, r1, r2, r3, r4}
}

func getRock(h int, mod int) rock {
	rocks := buildRocks()
	positions := [][2]int{
		{h + 4, 2},
		{h + 4, 3},
		{h + 4, 2},
		{h + 4, 2},
		{h + 4, 2}}
	return rockPos(rocks[mod], positions[mod])
}

func rockPos(r rock, p [2]int) rock {
	newR := rock{}
	for _, e := range r {
		newR = append(newR, [2]int{e[0] + p[0], e[1] + p[1]})
	}
	return newR
}

func moveRock(r rock, dir byte) rock {
	// dir values: < left, > right, default down
	switch dir {
	case '<':
		return rockPos(r, [2]int{0, -1})
	case '>':
		return rockPos(r, [2]int{0, 1})
	default:
		return rockPos(r, [2]int{-1, 0})
	}
}

func isPossible(r rock, prev map[[2]int]bool) bool {
	for _, e := range r {
		if e[0] < 0 || e[1] < 0 || e[1] > 6 || prev[e] {
			return false
		}
	}
	return true
}

func sum(s []int) int {
	r := 0
	for _, e := range s {
		r += e
	}
	return r
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()

	instructions := scn.Text()

	prev := map[[2]int]bool{}
	highest := -1
	dirCount := 0
	h := []int{}

	for j := 0; j < 4975; j++ {
		stopped := false
		r := getRock(highest, j%5)
		for !stopped {
			dir := instructions[dirCount%(len(instructions))]
			dirCount++
			moved := moveRock(r, dir)
			if isPossible(moved, prev) {
				r = moved
			}

			down := moveRock(r, 'x')
			if !isPossible(down, prev) {
				stopped = true
				for _, e := range r {
					highest = max(highest, e[0])
					prev[e] = true
				}
			} else {
				r = down
			}
		}
		h = append(h, highest)
	}

	fmt.Println(highest)
	fmt.Println()

	diffs := []int{}

	for i := 0; i < len(h)-1; i++ {
		diffs = append(diffs, h[i+1]-h[i])
	}

	c := []int{3, 3, 0, 0, 1, 3, 0, 3, 2, 0, 0, 2, 0, 0}
	// c := []int{3, 2, 0, 0, 1, 3, 3, 4, 0, 1, 2}

	for i := 0; i < len(diffs)-len(c); i++ {
		if slices.Equal(diffs[i:i+len(c)], c) {
			fmt.Println(i)
		}
	}

	fmt.Println()
	fmt.Println(h[2800] - h[1555])
	// fmt.Println(h[4910] - h[4875])
}
