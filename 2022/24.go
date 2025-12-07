//go:build ignore

package main

import (
	"fmt"
	"os"
	"strings"
)

type pair[T any] struct {
	data     T
	distance int
}

func newPair[T any](data T, distance int) pair[T] {
	return pair[T]{data: data, distance: distance}
}

type MinHeap[T any] struct {
	data []T
	less func(a, b T) bool
}

func NewMinHeap[T any](less func(a, b T) bool) *MinHeap[T] {
	return &MinHeap[T]{data: []T{}, less: less}
}

func (h *MinHeap[T]) Push(value T) {
	h.data = append(h.data, value)
	h.up(h.Len() - 1)
}

func (h *MinHeap[T]) Pop() T {
	if h.Len() == 0 {
		panic("empty heap")
	}
	n := h.Len() - 1
	h.data[0], h.data[n] = h.data[n], h.data[0]
	h.down(0, n)

	val := h.data[n]
	h.data = h.data[:n]
	return val
}

func (h *MinHeap[T]) up(j int) {
	for j > 0 {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(h.data[j], h.data[i]) {
			break
		}
		h.data[j], h.data[i] = h.data[i], h.data[j]
		j = i
	}
}

func (h *MinHeap[T]) down(i0, n int) {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(h.data[j2], h.data[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(h.data[j], h.data[i]) {
			break
		}
		h.data[i], h.data[j] = h.data[j], h.data[i]
		i = j
	}
}

func (h *MinHeap[T]) Len() int {
	return len(h.data)
}

var R int
var C int

type pt [2]int

type board map[pt][]byte

var boards []board

// boardIdx, posR, posC
type state [3]int

func (b board) copy() board {
	bb := board{}
	for k, v := range b {
		bb[k] = v
	}
	return bb
}

func (b board) String() string {
	s := ""
	for r := range R {
		for c := range C {
			e := b[pt{r, c}]
			switch len(e) {
			case 0:
				s += "."
			case 1:
				s += fmt.Sprintf("%c", e[0])
			default:
				s += fmt.Sprintf("%v", len(e))
			}
		}
		s += "\n"
	}
	return s
}

func (b board) pass() board {
	next := board{}
	for r := range R {
		for c := range C {
			for _, e := range b[pt{r, c}] {
				switch e {
				case '>':
					cc := c%(C-2) + 1
					next[pt{r, cc}] = append(next[pt{r, cc}], '>')
				case '<':
					cc := (c-2)%(C-2) + 1
					if cc < 1 {
						ccc := c + (C - 2)
						cc = (ccc-2)%(C-2) + 1
					}
					next[pt{r, cc}] = append(next[pt{r, cc}], '<')
				case '^':
					rr := (r-2)%(R-2) + 1
					if rr < 1 {
						rrr := r + (R - 2)
						rr = (rrr-2)%(R-2) + 1
					}
					next[pt{rr, c}] = append(next[pt{rr, c}], '^')
				case 'v':
					next[pt{r%(R-2) + 1, c}] = append(next[pt{r%(R-2) + 1, c}], 'v')
				case '#':
					next[pt{r, c}] = []byte{'#'}
				}
			}
		}
	}
	return next
}

func (b board) hash() int {
	hash := 0
	for r := range R {
		for c := range C {
			for _, e := range b[pt{r, c}] {
				hash = (hash*31 + int(e)) % 1_000_000_007
			}
		}
	}
	return hash
}

func (b board) end() bool {
	return len(b[pt{R - 1, C - 2}]) == 1 && b[pt{R - 1, C - 2}][0] == 'E'
}

func main() {
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))
	lines := strings.Split(input, "\n")

	R, C = len(lines), len(lines[0])

	diagram := board{}
	for r, row := range lines {
		for c := range len(row) {
			if row[c] == '.' {
				continue
			}
			diagram[pt{r, c}] = []byte{row[c]}
		}
	}

	hashes := map[int]bool{}

	for {
		h := diagram.hash()
		if v, ok := hashes[h]; ok && v {
			break
		}
		hashes[diagram.hash()] = true
		boards = append(boards, diagram)
		diagram = diagram.pass()
	}

	A := pt{0, 1}
	B := pt{R - 1, C - 2}

	// part1
	first, d0 := dijkstra(state{0, 0, 1}, B)
	fmt.Println(d0)

	// part2
	second, d1 := dijkstra(first, A)
	_, d2 := dijkstra(second, B)
	fmt.Println(d0 + d1 + d2)
}

func dijkstra(start state, end pt) (state, int) {
	dist := map[[3]int]int{}
	heap := NewMinHeap(func(a, b pair[state]) bool {
		return a.distance < b.distance
	})

	dist[start] = 0
	heap.Push(newPair(start, 0))

	for heap.Len() > 0 {
		u := heap.Pop()

		if (pt{u.data[1], u.data[2]}) == end {
			return u.data, u.distance
		}

		for _, v := range nextStates(u.data) {
			alt := u.distance + 1
			if dv, ok := dist[v]; !ok || alt < dv {
				dist[v] = alt
				heap.Push(newPair(v, alt))
			}
		}
	}
	panic("end unreached")
}

func nextStates(u state) []state {
	nextIdx, E := (u[0]+1)%len(boards), pt{u[1], u[2]}
	next := boards[nextIdx]
	var neighbors []state

	// wait
	if len(next[E]) == 0 {
		neighbors = append(neighbors, state{nextIdx, E[0], E[1]})
	}

	// move
	for _, dir := range []pt{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		EE := pt{E[0] + dir[0], E[1] + dir[1]}
		if 0 <= EE[0] && EE[0] < R && 0 <= EE[1] && EE[1] < C && len(next[EE]) == 0 {
			neighbors = append(neighbors, state{nextIdx, EE[0], EE[1]})
		}
	}

	return neighbors
}
