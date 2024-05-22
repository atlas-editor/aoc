//go:build ignore

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

func sandFall(rocks map[[2]int]bool) int {
	start := [2]int{500, 0}
	lb := 0
	for k := range rocks {
		lb = max(lb, k[1])
	}
	lb += 2
	c := 0
	solid := rocks

	for {
		if solid[start] {
			break
		}
		curr := start
		for {
			b0, b1, b2 := [2]int{curr[0] - 1, curr[1] + 1}, [2]int{curr[0], curr[1] + 1}, [2]int{curr[0] + 1, curr[1] + 1}
			if solid[b1] {
				if !solid[b0] {
					curr = b0
				} else if !solid[b2] {
					curr = b2
				} else {
					solid[curr] = true
					break
				}
			} else if curr[1] == lb-1 {
				solid[curr] = true
				break
			} else {
				curr = b1
			}
		}
		c++
	}

	return c
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	rocks := map[[2]int]bool{}
	for scn.Scan() {
		inputLine := scn.Text()
		lines := strings.Split(inputLine, " -> ")

		for i := 1; i < len(lines); i++ {
			d0, d1 := lines[i-1], lines[i]
			prev := slicesMap(strings.Split(d0, ","), atoi)
			next := slicesMap(strings.Split(d1, ","), atoi)

			if prev[0] == next[0] {
				for j := min(prev[1], next[1]); j <= max(prev[1], next[1]); j++ {
					rocks[[2]int{prev[0], j}] = true
				}
			} else if prev[1] == next[1] {
				for j := min(prev[0], next[0]); j <= max(prev[0], next[0]); j++ {
					rocks[[2]int{j, prev[1]}] = true
				}
			}
		}
	}

	fmt.Println(sandFall(rocks))
}
