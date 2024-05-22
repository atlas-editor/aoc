package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func slicesMap[S, T any](ts []S, f func(S) T) []T {
	us := []T{}
	for _, e := range ts {
		us = append(us, f(e))
	}
	return us
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func isVisible(grid [][]int, r, c int) int {
	R := len(grid)
	C := len(grid[0])

	// if r == 0 || r == R-1 || c == 0 || c == C-1 {
	// 	return true
	// }

	h := grid[r][c]

	up := 0
	for i := r - 1; i >= 0; i-- {
		up++
		if grid[i][c] >= h {
			break
		}
	}

	left := 0
	for i := c - 1; i >= 0; i-- {
		left++
		if grid[r][i] >= h {
			break
		}
	}

	down := 0
	for i := r + 1; i < R; i++ {
		down++
		if grid[i][c] >= h {
			break
		}
	}

	right := 0
	for i := c + 1; i < C; i++ {
		right++
		if grid[r][i] >= h {
			break
		}
	}

	return up * left * down * right
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	grid := [][]int{}
	for scanner.Scan() {
		sl := scanner.Text()
		line := slicesMap(strings.Split(sl, ""), atoi)
		grid = append(grid, line)
	}
	R := len(grid)
	C := len(grid[0])

	sum := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			sum = max(sum, isVisible(grid, i, j))
		}
	}
	fmt.Println(sum)

}
