package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func counter(h []string) map[string]int {
	unique := map[string]int{}
	for _, r := range h {
		unique[r]++
	}
	return unique
}

func sum(m [][]string) int {
	res := 0
	for i := 0; i < len(m); i++ {
		res += counter(m[i])["O"] * (len(m) - i)
	}
	return res
}

func findPlace(m [][]string, dir, i, j int) (int, int) {
	if dir == 0 {
		next := i - 1
		for next >= 0 && m[next][j] == "." {
			next--
		}

		if next == -1 {
			return 0, j
		}

		return next + 1, j
	}

	if dir == 1 {
		next := j - 1
		for next >= 0 && m[i][next] == "." {
			next--
		}

		if next == -1 {
			return i, 0
		}

		return i, next + 1
	}

	if dir == 2 {
		next := i + 1
		for next < len(m) && m[next][j] == "." {
			next++
		}

		if next == len(m) {
			return len(m) - 1, j
		}

		return next - 1, j
	}

	if dir == 3 {
		next := j + 1
		for next < len(m[0]) && m[i][next] == "." {
			next++
		}

		if next == len(m[0]) {
			return i, len(m[0]) - 1
		}

		return i, next - 1
	}

	return -1, -1
}

func tilt(m [][]string, dir int) [][]string {
	if dir == 0 || dir == 1 {
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[0]); j++ {
				if m[i][j] == "O" {
					newI, newJ := findPlace(m, dir, i, j)
					m[i][j] = "."
					m[newI][newJ] = "O"
				}
			}
		}
	} else {
		for i := len(m) - 1; i >= 0; i-- {
			for j := len(m[0]) - 1; j >= 0; j-- {
				if m[i][j] == "O" {
					newI, newJ := findPlace(m, dir, i, j)
					m[i][j] = "."
					m[newI][newJ] = "O"
				}
			}
		}
	}
	return m
}

func lastIndex(s []int, n int) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == n {
			return i
		}
	}
	return -1
}

// this is not rigirous but eh
func hasPeriodicTail(s []int, idx int) bool {
	l := s[idx:]
	period := len(s) - idx
	ll := s[idx-period : idx]
	return slices.Equal(l, ll)
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	m := [][]string{}
	for scn.Scan() {
		line := scn.Text()
		l := strings.Split(line, "")
		m = append(m, l)
	}
	sums := []int{}
	for i := 0; i < 1000000000; i++ {
		for j := 0; j < 4; j++ {
			m = tilt(m, j)
		}
		curr := sum(m)
		if li := lastIndex(sums, curr); li != -1 && hasPeriodicTail(sums, li) {
			idx := (1000000000 - i - 1) % len(sums[li:])
			fmt.Println(sums[li:][idx])
			return
		}
		sums = append(sums, curr)
	}
}
