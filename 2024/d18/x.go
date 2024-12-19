package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const R = 71
const C = 71
const LIMIT = 1024

func main() {
	path := os.Args[1]

	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(input string) int {
	m := parse(input)
	d, _ := bfs(m)
	return d
}

func p2(input string) string {
	m := parse(input)

	for _, line := range strings.Split(input, "\n")[LIMIT:] {
		coords := ints(line)
		r, c := coords[1], coords[0]
		m[r][c] = '#'
		if _, ok := bfs(m); !ok {
			return line
		}
	}

	panic("unreachable")
}

func parse(input string) [R][C]byte {
	m := [R][C]byte{}

	for _, line := range strings.Split(input, "\n")[:LIMIT] {
		coords := ints(line)
		r, c := coords[1], coords[0]
		m[r][c] = '#'
	}

	return m
}

func nbrs4(r, c int) []pt {
	n := []pt{}
	for _, d := range []pt{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		dr, dc := d[0], d[1]
		rr, cc := r+dr, c+dc
		if 0 <= rr && rr < R && 0 <= cc && cc < C {
			n = append(n, pt{rr, cc})
		}
	}
	return n
}

func bfs(m [R][C]byte) (int, bool) {
	start := pt{0, 0}
	end := pt{R - 1, C - 1}

	dist := map[pt]int{start: 0}

	q := []pt{start}
	seen := set[pt]{start: true}
	for len(q) > 0 {
		curr := popFront(&q)
		if curr == end {
			return dist[curr], true
		}

		for _, n := range nbrs4(curr[0], curr[1]) {
			if !seen[n] && m[n[0]][n[1]] != '#' {
				q = append(q, n)
				dist[n] = dist[curr] + 1
				seen[n] = true
			}
		}
	}

	return -1, false
}

/*
utils
*/

type set[T comparable] map[T]bool

type pt [2]int

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func ints(s string) (r []int) {
	p := regexp.MustCompile(`-?\d+`)
	for _, e := range p.FindAllString(s, -1) {
		r = append(r, atoi(e))
	}
	return
}

func popFront[T any](slice *[]T) T {
	if len(*slice) == 0 {
		panic("empty slice")
	}
	front := (*slice)[0]
	*slice = (*slice)[1:]
	return front
}
