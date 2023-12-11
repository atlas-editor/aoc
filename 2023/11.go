package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type pair struct {
	x, y int
}

func transpose(m [][]string) [][]string {
	rows := len(m)
	cols := len(m[0])

	t := make([][]string, cols)

	for i := 0; i < cols; i++ {
		t[i] = make([]string, rows)
		for j := 0; j < rows; j++ {
			t[i][j] = m[j][i]
		}
	}

	return t
}

func isConst(ss []string) bool {
	for _, r := range ss {
		if r != "." {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	xConst := []int{}
	yConst := []int{}

	u := [][]string{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()

		row := []string{}
		for _, r := range line {
			row = append(row, string(r))
		}
		u = append(u, row)
		if isConst(row) {
			xConst = append(xConst, i)
		}
		i++
	}

	uu := transpose(u)
	for i, s := range uu {
		if isConst(s) {
			yConst = append(yConst, i)
		}
	}

	gs := []pair{}
	for i, p := range u {
		for j, r := range p {
			if r == "#" {
				emptyX, _ := slices.BinarySearch(xConst, i)
				emptyY, _ := slices.BinarySearch(yConst, j)
				gs = append(gs, pair{i + (emptyX * (1000000 - 1)), j + (emptyY * (1000000 - 1))})
			}
		}
	}

	sum := 0
	for i := 0; i < len(gs); i++ {
		for j := i + 1; j < len(gs); j++ {
			x := gs[i].x - gs[j].x
			if x < 0 {
				x = -x
			}
			y := gs[i].y - gs[j].y
			if y < 0 {
				y = -y
			}
			sum += x + y
		}
	}
	fmt.Println(sum)
}
