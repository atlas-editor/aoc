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
	d, _ := dijkstra(input)
	return d
}

func p2(input string) int {
	_, s := dijkstra(input)
	return s
}

func dijkstra(input string) (int, int) {
	m := readMatrix(input, func(b byte) byte {
		return b
	})

	R, C := len(m), len(m[0])
	start := vec{}
	end := vec{}
	for r := range R {
		for c := range C {
			if m[r][c] == 'S' {
				start = vec{r, c}
			}
			if m[r][c] == 'E' {
				end = vec{r, c}
			}
		}
	}

	startState := state{start, vec{0, 1}}
	dist := map[state]int{startState: 0}
	seen := set[item[state]]{}
	prev := map[state]set[state]{startState: {}}
	h := newMinHeap[item[state]]()
	h.push(newItem(startState, 0))
	for h.len() > 0 {
		curr := h.pop()
		if seen[*curr] {
			continue
		}
		seen[*curr] = true

		s, distance := curr.state, curr.distance

		if s.pos == end {
			bestSpots := set[vec]{}
			var f func(state)
			f = func(s state) {
				bestSpots[s.pos] = true
				for s2 := range prev[s] {
					f(s2)
				}
			}
			f(s)
			return distance, len(bestSpots)
		}

		for _, n := range getNeighbors(curr.state, curr.distance, m) {
			nbr, alt := n.s, n.d
			if v, ok := dist[nbr]; !ok || alt <= v {
				if alt < v || !ok {
					prev[nbr] = set[state]{}
				}
				prev[nbr][s] = true
				dist[nbr] = alt
				h.push(newItem(nbr, alt))
			}
		}
	}

	panic("path not found")
}

type state struct {
	pos vec
	dir vec
}

type neighbor struct {
	s state
	d int
}

func getNeighbors(s state, d int, m [][]byte) []neighbor {
	R, C := len(m), len(m[0])
	neighbors := []neighbor{}
	forward := s.pos.add(s.dir)
	if 0 <= forward[0] && forward[0] < R && 0 <= forward[1] && forward[1] < C && m[forward[0]][forward[1]] != '#' {
		neighbors = append(neighbors, neighbor{state{forward, s.dir}, d + 1})
	}
	for _, ddir := range []vec{s.dir.rotate(1), s.dir.rotate(3)} {
		neighbors = append(neighbors, neighbor{state{s.pos, ddir}, d + 1000})
	}

	return neighbors
}

/*
utils
*/

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

type lesser[T any] interface {
	less(other T) bool
}

type item[T any] struct {
	state    T
	distance int
}

func newItem[T any](state T, distance int) *item[T] {
	return &item[T]{state: state, distance: distance}
}

func (p item[T]) less(other item[T]) bool {
	return p.distance < other.distance
}

type minHeap[T lesser[T]] struct {
	arr []*T
}

func newMinHeap[T lesser[T]]() *minHeap[T] {
	return &minHeap[T]{arr: []*T{}}
}

func (h *minHeap[T]) push(value *T) {
	h.arr = append(h.arr, value)
	h.up(h.len() - 1)
}

func (h *minHeap[T]) pop() *T {
	if h.len() == 0 {
		panic("empty heap")
	}
	min_ := h.arr[0]
	n := h.len() - 1
	h.arr[0] = h.arr[n]
	h.arr = h.arr[:n]
	h.down(0, n)
	return min_
}

func (h *minHeap[T]) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || !(*h.arr[j]).less(*h.arr[i]) {
			break
		}
		h.arr[j], h.arr[i] = h.arr[i], h.arr[j]
		j = i
	}
}

func (h *minHeap[T]) down(i0, n int) {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && (*h.arr[j2]).less(*h.arr[j1]) {
			j = j2
		}
		if !(*h.arr[j]).less(*h.arr[i]) {
			break
		}
		h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
		i = j
	}
}

func (h *minHeap[T]) len() int {
	return len(h.arr)
}
