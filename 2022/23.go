//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

type pt [2]int

var dirs = []pt{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
var sides = [][]pt{{{-1, 0}, {-1, 1}, {-1, -1}}, {{1, 0}, {1, 1}, {1, -1}}, {{0, -1}, {-1, -1}, {1, -1}}, {{0, 1}, {1, 1}, {-1, 1}}}

func main() {
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	scan := map[pt]bool{}
	for r, row := range strings.Split(input, "\n") {
		for c, el := range row {
			if el == '#' {
				scan[pt{r, c}] = true
			}
		}
	}

	idx := 0
	for {
		proposed := map[pt][]pt{}
		moved := false

		for p, val := range scan {
			if !val {
				continue
			}
			r, c := p[0], p[1]
			if rr, cc, ok := step1(r, c, scan, idx); ok {
				to := pt{rr, cc}
				proposed[to] = append(proposed[to], pt{r, c})
			}
		}

		scan, moved = step2(scan, proposed)
		if !moved {
			fmt.Println(idx + 1)
			return
		}
		idx++
	}
}

//func emptyTiles(scan [][]byte) int {
//	on := 0
//	minR, maxR, minC, maxC := 1<<32, -1<<32, 1<<32, -1<<32
//	for r, row := range scan {
//		for c, b := range row {
//			if b == '#' {
//				minR = min(minR, r)
//				maxR = max(maxR, r)
//				minC = min(minC, c)
//				maxC = max(maxC, c)
//				on++
//			}
//		}
//	}
//
//	return (maxR-minR+1)*(maxC-minC+1) - on
//}
//
//func extend(scan [][]byte, f int) [][]byte {
//	C := Len(scan[0])
//	empty := []byte{}
//	for range C {
//		empty = append(empty, '.')
//	}
//
//	for range f {
//		scan = append(scan, empty)
//		scan = append([][]byte{empty}, scan...)
//	}
//
//	for r := range Len(scan) {
//		for range f {
//			scan[r] = append(scan[r], '.')
//			scan[r] = append([]byte{'.'}, scan[r]...)
//		}
//	}
//
//	return scan
//}

func printscan(scan [][]byte) {
	for _, row := range scan {
		for _, b := range row {
			fmt.Printf("%c", b)
		}
		fmt.Println()
	}
}

func inBound(r, c, R, C int) bool {
	return 0 <= r && r < R && 0 <= c && c < C
}

func step1(r, c int, scan map[pt]bool, offset int) (int, int, bool) {
	m := true
outer:
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if i == 0 && j == 0 {
				continue
			}
			if scan[pt{r + i, c + j}] {
				m = false
				break outer
			}
		}
	}

	if m {
		return -1, -1, false
	}

outer2:
	for i := range 4 {
		ii := (i + offset) % 4
		for _, s := range sides[ii] {
			if scan[pt{r + s[0], c + s[1]}] {
				continue outer2
			}
		}
		return r + dirs[ii][0], c + dirs[ii][1], true
	}

	return -1, -1, false
}

func step2(scan map[pt]bool, proposed map[pt][]pt) (map[pt]bool, bool) {
	moved := false
	for to, from := range proposed {
		if len(from) != 1 {
			continue
		}
		toPt := pt{to[0], to[1]}
		fromPt := pt{from[0][0], from[0][1]}
		scan[toPt] = true
		scan[fromPt] = false
		moved = true
	}
	return scan, moved
}
