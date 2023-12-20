package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	s    state
	Cost int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Cost < pq[j].Cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type pair [2]int

type tuple struct {
	r, c, blockLen int
	dir            pair
}

type state struct {
	pos, dir pair
	run      int
}

var north pair = pair{-1, 0}
var west pair = pair{0, -1}
var south pair = pair{1, 0}
var east pair = pair{0, 1}

var dirs []pair = []pair{north, west, south, east}

var cache = make(map[tuple]int)
var visited = make(map[tuple]bool)

func isInside(r, c int, maze [][]int) bool {
	return r >= 0 && r < len(maze) && c >= 0 && c < len(maze[0])
}

func isInsideFour(r, c int, dir pair, maze [][]int) bool {
	return r+(4*dir[0]) >= 0 && r+(4*dir[0]) < len(maze) && c+(4*dir[1]) >= 0 && c+(4*dir[1]) < len(maze[0])
}

func isPossible(r, c, blockLen int, dir pair, maze [][]int) bool {
	for i := 0; i < blockLen; i++ {
		r -= dir[0]
		c -= dir[1]
	}

	return isInside(r, c, maze)
}

func dirChange(d pair) []pair {
	if d == north || d == south {
		return []pair{west, east}
	}

	if d == west || d == east {
		return []pair{north, south}
	}

	return []pair{}
}

func nbrs(s state, maze [][]int) []state {
	res := []state{}
	nextPos := pair{s.pos[0] + s.dir[0], s.pos[1] + s.dir[1]}

	if s.run < 10 && isInside(nextPos[0]+s.dir[0], nextPos[1]+s.dir[1], maze) {
		ns := state{nextPos, s.dir, s.run + 1}
		if s.run < 4 {
			return []state{ns}
		}
		res = append(res, ns)
	}

	for _, d := range dirChange(s.dir) {
		if isInsideFour(nextPos[0], nextPos[1], d, maze) {
			ns := state{nextPos, d, 1}
			res = append(res, ns)
		}
	}
	return res
}

func djikstra(maze [][]int) int {
	R, C := len(maze), len(maze[0])

	startState := state{pair{0, 0}, east, 1}
	targetPos := pair{R - 1, C - 1}

	dist := map[state]int{}
	prev := map[state]state{}
	pQ := make(PriorityQueue, 0)
	heap.Init(&pQ)

	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			for _, d := range dirs {
				for i := 0; i < 11; i++ {
					curr := state{pair{r, c}, d, i}
					dist[curr] = 1e9
					prev[curr] = state{pair{-1, -1}, d, i}
				}
			}
		}
	}

	dist[startState] = 0
	heap.Push(&pQ, &Node{startState, 0})
	for pQ.Len() > 0 {
		uu := *heap.Pop(&pQ).(*Node)
		u := uu.s
		if u.pos == targetPos {
			return dist[u]
		}

		for _, v := range nbrs(u, maze) {
			alt := dist[u] + maze[v.pos[0]][v.pos[1]]
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				heap.Push(&pQ, &Node{v, alt})
			}

		}
	}

	return -1
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	m := [][]int{}
	for scn.Scan() {
		line := scn.Text()
		numsStr := strings.Split(line, "")
		nums := []int{}
		for _, v := range numsStr {
			w, _ := strconv.Atoi(v)
			nums = append(nums, w)
		}
		m = append(m, nums)
	}

	r := djikstra(m)
	fmt.Println(r)
}
