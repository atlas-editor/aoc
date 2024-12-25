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
	parts := strings.Split(input, "\n\n")

	locks := [][]int{}
	keys := [][]int{}
	for _, p := range parts {
		m := readMatrix(p, func(b byte) byte {
			return b
		})

		R, C := len(m), len(m[0])

		if isLock(m) {
			lock := []int{}
			for c := range C {
				for r := range R {
					if m[r][c] != '#' {
						lock = append(lock, r-1)
						break
					}
				}
			}
			locks = append(locks, lock)
		} else {
			key := []int{}
			for c := range C {
				for r := range R {
					if m[r][c] != '.' {
						key = append(key, R-r-1)
						break
					}
				}
			}
			keys = append(keys, key)
		}
	}

	s := 0
	for _, lock := range locks {
		for _, key := range keys {
			if fit(lock, key) {
				s++
			}
		}
	}

	return s
}

func p2(input string) any {
	return nil
}

func isLock(m [][]byte) bool {
	R, C := len(m), len(m[0])
	for c := range C {
		if m[0][c] != '#' || m[R-1][c] != '.' {
			return false
		}
	}
	return true
}

func fit(lock, key []int) bool {
	for i := range len(lock) {
		if lock[i]+key[i] >= 6 {
			return false
		}
	}
	return true
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
