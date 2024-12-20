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
	return solve(input, 2)
}

func p2(input string) int {
	return solve(input, 20)
}

func solve(input string, cheatTime int) int {
	m := readMatrix(input, func(b byte) byte {
		return b
	})

	p := shortestPath(m)
	dte := map[pt]int{}
	for i, v := range p {
		dte[v] = len(p) - i - 1
	}

	s := 0
	for dts, u := range p {
		for _, v := range cheat(m, u, cheatTime) {
			if len(p)-1-(dts+dist(u, v)+dte[v]) >= 100 {
				s++
			}
		}
	}

	return s
}

func nbrs4(r, c, R, C int) []pt {
	n := []pt{}
	for _, d := range []pt{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		rr, cc := r+d[0], c+d[1]
		if 0 <= rr && rr < R && 0 <= cc && cc < C {
			n = append(n, pt{rr, cc})
		}
	}
	return n
}

func cheat(m [][]byte, pos pt, cheatTime int) []pt {
	R, C := len(m), len(m[0])
	posR, posC := pos[0], pos[1]

	n := []pt{}
	for dr := -cheatTime; dr <= cheatTime; dr++ {
		for dc := -abs(cheatTime - abs(dr)); dc <= abs(cheatTime-abs(dr)); dc++ {
			rr, cc := posR+dr, posC+dc
			if (posR != rr || posC != cc) && 0 <= rr && rr < R && 0 <= cc && cc < C && m[rr][cc] != '#' {
				n = append(n, pt{rr, cc})
			}
		}
	}
	return n
}

func shortestPath(m [][]byte) []pt {
	R, C := len(m), len(m[0])
	start := pt{}
	end := pt{}
	for r := range R {
		for c := range C {
			switch m[r][c] {
			case 'S':
				start = pt{r, c}
			case 'E':
				end = pt{r, c}

			}
		}
	}

	p := []pt{}
	q := []pt{start}
	seen := set[pt]{start: true}
	for len(q) > 0 {
		curr := pop(&q)
		p = append(p, curr)
		if curr == end {
			return p
		}

		for _, n := range nbrs4(curr[0], curr[1], R, C) {
			if !seen[n] && m[n[0]][n[1]] != '#' {
				q = append(q, n)
				seen[n] = true
			}
		}
	}

	panic("end unreached")
}

func dist(u, v pt) int {
	return abs(u[0]-v[0]) + abs(u[1]-v[1])
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
