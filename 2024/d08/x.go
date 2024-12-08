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
	return solve(input, true)
}

func p2(input string) int {
	return solve(input, false)
}

func solve(input string, twoAntinodes bool) int {
	m := readMatrix(input, func(b byte) byte {
		return b
	})

	R, C := len(m), len(m[0])

	freq := map[byte][]vec{}
	for r := range R {
		for c := range C {
			if m[r][c] == '.' {
				continue
			}
			freq[m[r][c]] = append(freq[m[r][c]], vec{r, c})
		}
	}

	antinodes := set[vec]{}
	for _, loc := range freq {
		for i := 0; i < len(loc); i++ {
			for j := i + 1; j < len(loc); j++ {
				x, y := loc[i], loc[j]
				diff := vec{x[0] - y[0], x[1] - y[1]}

				if twoAntinodes {
					addAntinodes([]vec{x.add(diff), y.add(diff.mul(-1))}, antinodes, R, C)
					continue
				}

				d := vec{0, 0}
				for {
					if !addAntinodes([]vec{x.add(d), x.add(d.mul(-1))}, antinodes, R, C) {
						break
					} else {
						d = d.add(diff)
					}
				}
			}
		}
	}

	return len(antinodes)
}

func addAntinodes(antinodes []vec, antinodeSet set[vec], R, C int) bool {
	inBound := false
	for _, a := range antinodes {
		if 0 <= a[0] && a[0] < R && 0 <= a[1] && a[1] < C {
			inBound = true
			antinodeSet[a] = true
		}
	}
	return inBound
}

/*
utils
*/

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

type vec [2]int

func (u vec) add(v vec) vec {
	return vec{u[0] + v[0], u[1] + v[1]}
}

func (u vec) mul(c int) vec {
	return vec{c * u[0], c * u[1]}
}

type set[T comparable] map[T]bool
