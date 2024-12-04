package main

import (
	"fmt"
	"os"
	"slices"
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
	m := readMatrix(input, func(b byte) byte {
		return b
	})

	s := 0
	R, C := len(m), len(m[0])
	for r := range len(m) {
		for c := range len(m[0]) {
		outer:
			for _, d := range []vec{{-1, 1}, {-1, -1}, {1, -1}, {1, 1}, {-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
				curr := vec{r, c}
				for i := range 4 {
					rr, cc := curr[0], curr[1]
					if rr < 0 || rr >= R || cc < 0 || cc >= C || m[rr][cc] != "XMAS"[i] {
						continue outer
					}
					curr = curr.add(d)
				}
				s++
			}
		}
	}

	return s
}

func p2(input string) int {
	m := readMatrix(input, func(b byte) byte {
		return b
	})

	s := 0
	R, C := len(m), len(m[0])
	for r := range R {
		for c := range C {
			if r < 1 || r > R-2 || c < 1 || c > C-2 || m[r][c] != 'A' {
				continue
			}
			ul := m[r-1][c-1]
			ur := m[r-1][c+1]
			dl := m[r+1][c-1]
			dr := m[r+1][c+1]
			switch {
			case ul == 'M' && ur == 'S' && dl == 'M' && dr == 'S',
				ul == 'M' && ur == 'M' && dl == 'S' && dr == 'S',
				ul == 'S' && ur == 'S' && dl == 'M' && dr == 'M',
				ul == 'S' && ur == 'M' && dl == 'S' && dr == 'M':
				s++
			}
		}
	}

	return s
}

/*
utils
*/

type vec [2]int

func (u vec) add(v vec) vec {
	return vec{u[0] + v[0], u[1] + v[1]}
}

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

func transpose[T any](m [][]T) [][]T {
	R, C := len(m), len(m[0])
	m2 := make([][]T, C)
	for i := range C {
		m2[i] = make([]T, R)
	}

	for r := range R {
		for c := range C {
			m2[c][r] = m[r][c]
		}
	}

	return m2
}

func rotate[T any](m [][]T, n int) [][]T {
	for range n % 4 {
		m = transpose(m)

		for c := range len(m) {
			slices.Reverse(m[c])
		}
	}
	return m
}
