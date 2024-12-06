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
	R, C := len(m), len(m[0])
outer:
	for r := range R {
		for c := range C {
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

func onLoop(m [][]byte, start, obstruction vec) (bool, map[vec]bool) {
	seenPt := map[vec]bool{}
	seenState := map[state]bool{}
	curr := start
	dir := vec{-1, 0}
	for {
		if _, ok := seenState[state{curr, dir}]; ok {
			return true, map[vec]bool{}
		}

		seenPt[curr] = true
		seenState[state{curr, dir}] = true

		next := curr.add(dir)
		r, c := next[0], next[1]

		if !(0 <= r && r < len(m) && c >= 0 && c < len(m[0])) {
			return false, seenPt
		}

		if m[r][c] == '#' || next == obstruction {
			dir = dir.rotate(3)
		} else {
			curr = next
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
