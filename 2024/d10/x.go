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
	return solve(input, hike)
}

func p2(input string) int {
	return solve(input, hikeAll)
}

func solve(input string, hikeFunc func([][]int, pt) int) int {
	m := readMatrix(input, func(b byte) int {
		return int(b - 48)
	})

	s := 0
	for r := range len(m) {
		for c := range len(m[0]) {
			if m[r][c] == 0 {
				s += hikeFunc(m, pt{r, c})
			}
		}
	}

	return s
}

func hike(m [][]int, start pt) int {
	q := []pt{start}
	seen := set[pt]{start: true}
	s := 0
	for len(q) > 0 {
		curr := pop(&q)
		r, c := curr[0], curr[1]

		if m[r][c] == 9 {
			s++
			continue
		}

		for _, n := range nbrs4(m, r, c, len(m), len(m[0])) {
			if !seen[n] {
				q = append(q, n)
				seen[n] = true
			}
		}
	}

	return s
}

func hikeAll(m [][]int, start pt) int {
	R, C := len(m), len(m[0])

	var dfs func(pt, set[pt]) int

	dfs = func(u pt, seen set[pt]) int {
		if m[u[0]][u[1]] == 9 {
			return 1
		}

		seen[u] = true

		s := 0
		for _, n := range nbrs4(m, u[0], u[1], R, C) {
			if !seen[n] {
				s += dfs(n, seen)
			}
		}

		seen[u] = false
		return s
	}

	return dfs(start, set[pt]{})
}

func nbrs4(m [][]int, r, c, R, C int) []pt {
	n := []pt{}
	curr := m[r][c]
	for _, d := range []pt{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		rr, cc := r+d[0], c+d[1]

		if 0 <= rr && rr < R && 0 <= cc && cc < C && m[rr][cc] == curr+1 {
			n = append(n, pt{rr, cc})
		}
	}
	return n
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
