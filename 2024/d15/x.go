package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var dirs = map[byte]vec{'^': {-1, 0}, 'v': {1, 0}, '>': {0, 1}, '<': {0, -1}}

func main() {
	path := os.Args[1]

	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(input string) int {
	parts := strings.Split(input, "\n\n")
	m := readMatrix(parts[0], func(b byte) byte {
		return b
	})
	R, C := len(m), len(m[0])

	moves := []byte{}
	for i := range parts[1] {
		if parts[1][i] != '\n' {
			moves = append(moves, parts[1][i])
		}
	}

	at := vec{}
outer:
	for r := range R {
		for c := range C {
			if m[r][c] == '@' {
				at = vec{r, c}
				break outer
			}
		}
	}

	for _, move := range moves {
		at = moveSimple(at, dirs[move], m)
	}

	gps := 0
	for r := range R {
		for c := range C {
			if m[r][c] == 'O' {
				gps += 100*r + c
			}
		}
	}

	return gps
}

func p2(input string) int {
	parts := strings.Split(input, "\n\n")
	small := readMatrix(parts[0], func(b byte) byte {
		return b
	})

	moves := []byte{}
	for i := range parts[1] {
		if parts[1][i] != '\n' {
			moves = append(moves, parts[1][i])
		}
	}

	m := make([][]byte, len(small))
	for i := range m {
		m[i] = []byte{}
		for _, c := range small[i] {
			switch c {
			case '#':
				m[i] = append(m[i], "##"...)
			case 'O':
				m[i] = append(m[i], "[]"...)
			case '.':
				m[i] = append(m[i], ".."...)
			case '@':
				m[i] = append(m[i], "@."...)
			}
		}
	}

	R, C := len(m), len(m[0])

	at := vec{}
outer:
	for r := range R {
		for c := range C {
			if m[r][c] == '@' {
				at = vec{r, c}
				break outer
			}
		}
	}

	for _, move := range moves {
		switch move {
		case '<', '>':
			at = moveHorizontal(at, dirs[move], m)
		case '^', 'v':
			at = moveVertical(at, dirs[move], m)
		}
	}

	gps := 0
	for r := range R {
		for c := range C {
			if m[r][c] == '[' {
				gps += 100*r + c
			}
		}
	}

	return gps
}

func inBound(r, c, R, C int) bool {
	return 0 <= r && r < R && 0 <= c && c < C
}

func moveSimple(pos, dir vec, m [][]byte) vec {
	R, C := len(m), len(m[0])
	next := pos.add(dir)
	for {
		r, c := next[0], next[1]
		switch {
		case !inBound(r, c, R, C) || m[r][c] == '#':
			return pos
		case m[r][c] == '.':
			m[next[0]][next[1]] = 'O'
			m[pos[0]][pos[1]] = '.'
			pos = pos.add(dir)
			m[pos[0]][pos[1]] = '@'
			return pos
		}
		next = next.add(dir)
	}
}

func moveHorizontal(pos, dir vec, m [][]byte) vec {
	R, C := len(m), len(m[0])
	next := pos.add(dir)
	for {
		r, c := next[0], next[1]
		switch {
		case !inBound(r, c, R, C) || m[r][c] == '#':
			return pos
		case m[r][c] == '.':
			m[pos[0]][pos[1]] = '.'
			pos = pos.add(dir)
			m[pos[0]][pos[1]] = '@'

			if next != pos {
				row := pos[0]

				start := pos.add(dir)[1]
				end := next[1]
				for i := min(start, end); i <= max(start, end); i += 2 {
					m[row][i] = '['
					m[row][i+1] = ']'
				}
			}

			return pos
		}

		next = next.add(dir)
	}
}

func moveVertical(pos, dir vec, m [][]byte) vec {
	next := pos.add(dir)
	if m[next[0]][next[1]] == '#' {
		return pos
	}

	blocks := blocksToMove(pos, m, dir)
	// check if all blocks in this group can be moved in the given direction
	for _, level := range blocks {
		for block := range level {
			moveTo := block.add(dir)
			if m[moveTo[0]][moveTo[1]] == '#' {
				return pos
			}
		}
	}

	// each block is moveable, and they are sorted from furthest to closes to your avatar @, hence we just move them level by level, block by block
	for _, level := range blocks {
		for block := range level {
			nextTo := block.add(dir)
			curr := m[block[0]][block[1]]
			m[nextTo[0]][nextTo[1]] = curr
			m[block[0]][block[1]] = '.'
		}
	}

	m[pos[0]][pos[1]] = '.'
	pos = pos.add(dir)
	m[pos[0]][pos[1]] = '@'

	return pos
}

// find all blocks that are stacked onto each other in the given direction
func blocksToMove(pos vec, m [][]byte, dir vec) []set[vec] {
	R, C := len(m), len(m[0])
	blockPositions := []set[vec]{}
	for range R {
		blockPositions = append(blockPositions, set[vec]{})
	}

	var find func(vec, int)
	find = func(v vec, lvl int) {
		next := v.add(dir)

		nextLevel := []vec{}
		r, c := next[0], next[1]

		if !inBound(r, c, R, C) {
			return
		}

		switch m[r][c] {
		case '[':
			left, right := next, next.add(vec{0, 1})
			nextLevel = append(nextLevel, []vec{left, right}...)

			blockPositions[lvl][left] = true
			blockPositions[lvl][right] = true
		case ']':
			left, right := next.add(vec{0, -1}), next
			nextLevel = append(nextLevel, []vec{left, right}...)

			blockPositions[lvl][left] = true
			blockPositions[lvl][right] = true
		}

		for _, n := range nextLevel {
			find(n, lvl+1)
		}
	}

	find(pos, 0)

	i := 0
	for len(blockPositions[i]) != 0 {
		i++
	}
	blockPositions = blockPositions[:i]
	slices.Reverse(blockPositions)

	return blockPositions
}

/*
utils
*/

type vec [2]int

func (u vec) add(v vec) vec {
	return vec{u[0] + v[0], u[1] + v[1]}
}

type set[T comparable] map[T]bool

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
