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
	area, perimeter := fence(readMatrix(input, func(b byte) byte {
		return b
	}))

	r := 0
	for i := range len(area) {
		p := 0
		for _, s := range perimeter[i] {
			p += s
		}
		r += area[i] * p
	}

	return r
}

func p2(input string) int {
	area, perimeter := fence(readMatrix(input, func(b byte) byte {
		return b
	}))

	r := 0
	for i := range len(area) {
		r += area[i] * len(perimeter[i])
	}

	return r
}

func fence(m [][]byte) ([]int, [][]int) {
	area := []int{}
	perimeter := [][]int{}

	seen := set[pt]{}
	R, C := len(m), len(m[0])

	for r := range R {
		for c := range C {
			if seen[pt{r, c}] {
				continue
			}

			regionId := m[r][c]
			currArea := 0
			seen[pt{r, c}] = true
			regionPts := []pt{{r, c}}
			q := []pt{{r, c}}

			for len(q) > 0 {
				curr := pop(&q)
				currArea++

				for _, n := range nbrs4(curr[0], curr[1], R, C) {
					if !seen[n] && m[n[0]][n[1]] == regionId {
						q = append(q, n)
						seen[n] = true
						regionPts = append(regionPts, n)
					}
				}
			}

			area = append(area, currArea)
			lengths := sideLengths(m, extractEdges(m, regionPts))
			perimeter = append(perimeter, lengths)
		}
	}

	return area, perimeter
}

type edge set[pt]

func (e edge) overlap(f edge) (pt, int) {
	o := 0
	var p pt
	for pt0 := range e {
		for pt1 := range f {
			if pt0 == pt1 {
				p = pt0
				o++
			}
		}
	}
	return p, o
}

func (e edge) orientation() byte {
	edges := []pt{}
	for p := range e {
		edges = append(edges, p)
	}
	if edges[0][0] == edges[1][0] {
		return '-'
	} else {
		return '|'
	}
}

func match(e, f edge, m [][]byte) bool {
	R, C := len(m), len(m[0])

	eOrientation := e.orientation()
	overlap, count := e.overlap(f)
	switch count {
	case 0:
		return false
	case 2:
		return true
	}
	if eOrientation != f.orientation() {
		return false
	}

	nbrs := []byte{0, 0, 0, 0}
	for i, d := range []pt{{-1, -1}, {-1, 0}, {0, -1}, {0, 0}} {
		r, c := overlap[0]+d[0], overlap[1]+d[1]
		if inBound(r, c, R, C) {
			nbrs[i] = m[r][c]
		}
	}

	switch eOrientation {
	case '-':
		return nbrs[0] == nbrs[1] || nbrs[2] == nbrs[3]
	case '|':
		return nbrs[0] == nbrs[2] || nbrs[1] == nbrs[3]
	default:
		panic("invalid orientation")
	}
}

func extractEdges(m [][]byte, pts []pt) []edge {
	R, C := len(m), len(m[0])
	regionId := m[pts[0][0]][pts[0][1]]

	edges := []edge{}
	for _, p := range pts {
		r, c := p[0], p[1]
		for _, d := range []pt{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			rr, cc := r+d[0], c+d[1]
			if !(0 <= rr && rr < R && 0 <= cc && cc < C) || m[rr][cc] != regionId {
				switch d {
				case pt{-1, 0}:
					edges = append(edges, edge{{r, c}: true, {r, c + 1}: true})
				case pt{1, 0}:
					edges = append(edges, edge{{r + 1, c}: true, {r + 1, c + 1}: true})
				case pt{0, -1}:
					edges = append(edges, edge{{r, c}: true, {r + 1, c}: true})
				case pt{0, 1}:
					edges = append(edges, edge{{r, c + 1}: true, {r + 1, c + 1}: true})
				}
			}
		}
	}

	return edges
}

func sideLengths(m [][]byte, edges []edge) []int {
	seen := set[int]{}
	lengths := []int{}

	for i, e := range edges {
		if seen[i] {
			continue
		}

		seen[i] = true
		currLen := 0
		q := []edge{e}

		for len(q) > 0 {
			currLen++
			currE := pop(&q)

			for j, ee := range edges {
				if !seen[j] && match(currE, ee, m) {
					q = append(q, ee)
					seen[j] = true
				}
			}
		}
		lengths = append(lengths, currLen)
	}

	return lengths
}

func nbrs4(r, c, R, C int) []pt {
	nbrs := []pt{}
	for _, d := range []pt{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		rr, cc := r+d[0], c+d[1]
		if inBound(rr, cc, R, C) {
			nbrs = append(nbrs, pt{rr, cc})
		}
	}
	return nbrs
}

func inBound(r, c, R, C int) bool {
	return 0 <= r && r < R && 0 <= c && c < C
}

/*
utils
*/

type pt [2]int

type set[T comparable] map[T]bool

func readMatrix[T any](s string, transform func(byte) T) [][]T {
	rows := strings.Split(s, "\n")
	matrix := make([][]T, len(rows))

	for i, row := range rows {
		matrix[i] = make([]T, len(row))
		for j := range row {
			matrix[i][j] = transform(row[j])
		}
	}

	return matrix
}

func pop[T any](slice *[]T) T {
	n := len(*slice)
	if n == 0 {
		panic("empty slice")
	}
	back := (*slice)[n-1]
	*slice = (*slice)[:n-1]
	return back
}
