package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func bfs(hill [][]string, start, end [2]int) int {
	letters := "abcdefghijklmnopqrstuvwxyz"
	R, C := len(hill), len(hill[0])

	q := [][2]int{start}
	distance := map[[2]int]int{start: 0}

	for len(q) > 0 {
		curr := q[0]
		currVal := hill[curr[0]][curr[1]]
		currDis := distance[curr]
		q = q[1:]

		if curr == end {
			return currDis
		}

		dr := []int{-1, 1, 0, 0}
		dc := []int{0, 0, -1, 1}

		for i := 0; i < 4; i++ {
			next := [2]int{curr[0] + dr[i], curr[1] + dc[i]}
			if !(0 <= next[0] && next[0] < R && 0 <= next[1] && next[1] < C) {
				continue
			}
			if _, found := distance[next]; found {
				continue
			}
			nextVal := hill[next[0]][next[1]]
			if strings.Index(letters, currVal)-strings.Index(letters, nextVal) >= -1 {
				q = append(q, next)
				distance[next] = currDis + 1
			}
		}
	}

	return -1
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	hill := [][]string{}
	as := [][2]int{}
	end := [2]int{}
	row := 0
	for scn.Scan() {
		line := strings.Split(scn.Text(), "")
		for i := 0; i < len(line); i++ {
			if line[i] == "a" || line[i] == "S" {
				as = append(as, [2]int{row, i})
			}
		}
		if v := slices.Index(line, "S"); v != -1 {
			line[v] = "a"
		}
		if v := slices.Index(line, "E"); v != -1 {
			end[0] = row
			end[1] = v
			line[v] = "z"
		}
		hill = append(hill, line)
		row++
	}
	r := 1000000000
	for _, start := range as {
		b := bfs(hill, start, end)
		if b < 0 {
			continue
		}
		r = min(r, b)
	}
	fmt.Println(r)
}
