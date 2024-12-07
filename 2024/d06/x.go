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
	m, start := parse(input)

	_, path := onLoop(m, start, vec{-1, -1})

	return len(path)
}

func p2(input string) int {
	m, start := parse(input)

	_, orig := onLoop(m, start, vec{-1, -1})
	c := make(chan bool)
	for p := range orig {
		if p == start {
			continue
		}
		go func() {
			ok, _ := onLoop(m, start, p)
			c <- ok
		}()
	}

	loops := 0
	for range len(orig) - 1 {
		if <-c {
			loops++
		}
	}

	return loops

}

func parse(input string) ([][]byte, vec) {
	m := readMatrix(input, func(b byte) byte {
		return b
	})

	start := vec{0, 0}
outer:
	for r := range len(m) {
		for c := range len(m[0]) {
			if m[r][c] == '^' {
				start = vec{r, c}
				break outer
			}

		}
	}

	return m, start
}

type state struct {
	pos vec
	dir vec
}

func extractPts(states set[state]) set[vec] {
	pts := set[vec]{}
	for s := range states {
		pts[s.pos] = true
	}
	return pts
}

func onLoop(m [][]byte, start, obstruction vec) (bool, set[vec]) {
	seen := set[state]{}
	curr := state{start, vec{-1, 0}}
	for {
		if _, ok := seen[curr]; ok {
			return true, set[vec]{}
		}

		seen[curr] = true

		nextPos := curr.pos.add(curr.dir)
		r, c := nextPos[0], nextPos[1]

		if !(0 <= r && r < len(m) && 0 <= c && c < len(m[0])) {
			return false, extractPts(seen)
		}

		if m[r][c] == '#' || nextPos == obstruction {
			curr.dir = curr.dir.rotate(3)
		} else {
			curr.pos = nextPos
		}
	}

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

type set[T comparable] map[T]bool

type vec [2]int

func (u vec) add(v vec) vec {
	return vec{u[0] + v[0], u[1] + v[1]}
}

func (u vec) rotate(n int) vec {
	a, b := u[0], u[1]
	for range n % 4 {
		a, b = -b, a
	}
	return vec{a, b}
}
