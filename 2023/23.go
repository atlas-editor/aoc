package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type pair struct {
	r, c int
}

type tuple struct {
	p pair
	d int
}

func nbrs(r, c int, maze [][]string, visited map[pair]bool) []pair {
	curr := maze[r][c]

	if curr != "." {
		switch curr {
		case ">":
			return []pair{pair{r, c + 1}}
		case "v":
			return []pair{pair{r + 1, c}}
		}
	}

	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}
	res := []pair{}
	for i := 0; i < 4; i++ {
		rr, cc := r+dr[i], c+dc[i]
		if !isValid(rr, cc, maze, visited) ||
			(maze[rr][cc] == ">" && dc[i] == -1) &&
				(maze[rr][cc] == "v" && dr[i] == -1) {
			continue
		}
		res = append(res, pair{rr, cc})
	}
	return res
}

func isSink(r, c int, maze [][]string) bool {
	if r == 0 || c == 0 {
		return false
	}
	return maze[r-1][c] == "v" && maze[r][c-1] == ">"
}

func isPossible(p pair, maze [][]string, visited map[pair]bool) bool {
	r, c := p.r, p.c
	return !isSink(r, c, maze) || (visited[pair{r - 1, c}] && visited[pair{r, c - 1}])
}

func isValid(r, c int, maze [][]string, visited map[pair]bool) bool {
	rows := len(maze)
	cols := len(maze[0])
	return r >= 0 && r < rows && c >= 0 && c < cols && maze[r][c] != "#" && !visited[pair{r, c}]
}

func bfs(maze [][]string) int {
	visited := map[pair]bool{}

	queue := list.New()
	queue.PushBack(pair{0, 1})

	topOrder := []pair{}
	out := map[pair][]pair{}

	for queue.Len() > 0 {
		curr := queue.Front()
		for curr != nil {
			if isPossible(curr.Value.(pair), maze, visited) {
				break
			}
			curr = curr.Next()
		}
		queue.Remove(curr)
		currPt := curr.Value.(pair)

		if visited[currPt] {
			continue
		}
		visited[currPt] = true

		topOrder = append(topOrder, currPt)
		out[currPt] = nbrs(currPt.r, currPt.c, maze, visited)

		for _, n := range nbrs(currPt.r, currPt.c, maze, visited) {
			queue.PushBack(n)
		}
	}

	in := map[pair][]pair{}

	for vertex, outNbrs := range out {
		for _, on := range outNbrs {
			in[on] = append(in[on], vertex)
		}
	}

	distances := map[pair]int{pair{0, 1}: 0}
	for idx, v := range topOrder {
		if idx == 0 {
			continue
		}
		cd := -1

		for _, prevN := range in[v] {
			cd = max(cd, distances[prevN]+1)
		}

		distances[v] = cd
	}

	R, C := len(maze), len(maze[0])
	return distances[pair{R - 1, C - 2}]
}

func bfs2(G map[pair][]pair, start pair) []tuple {
	visited := map[pair]bool{}
	nbrs := []tuple{}

	queue := list.New()
	queue.PushBack(tuple{start, 0})
	visited[start] = true

	for queue.Len() > 0 {
		curr := queue.Front()
		queue.Remove(curr)
		currVal := curr.Value.(tuple)
		u := currVal.p
		d := currVal.d

		if len(G[u]) != 2 && u != start {
			nbrs = append(nbrs, tuple{u, d})
			continue
		}

		for _, v := range G[u] {
			if visited[v] {
				continue
			}
			queue.PushBack(tuple{v, d + 1})
			visited[v] = true
		}
	}

	return nbrs
}

func getNbrs(r, c int, m [][]string) []pair {
	R, C := len(m), len(m[0])
	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}
	res := []pair{}

	for i := 0; i < 4; i++ {
		rr, cc := r+dr[i], c+dc[i]
		if 0 <= rr && rr < R && 0 <= c && c < C && m[rr][cc] != "#" {
			res = append(res, pair{rr, cc})
		}
	}

	return res
}

func dfs(G map[pair][]tuple, start, end pair) int {
	visited := []tuple{}

	stack := list.New()
	stack.PushBack(G[start])
	visited = append(visited, tuple{start, 0})

	lp := 0
OuterLoop:
	for stack.Len() > 0 {
		curr := stack.Back()
		nbrs := curr.Value.([]tuple)
		if len(nbrs) == 0 {
			stack.Remove(curr)
			visited = visited[:len(visited)-1]
			continue
		}
		stack.Remove(curr)
		stack.PushBack(nbrs[1:])

		nbrVal := nbrs[0]
		u := nbrVal.p
		d := nbrVal.d

		for _, v := range visited {
			if u == v.p {
				continue OuterLoop
			}
		}

		visited = append(visited, tuple{u, visited[len(visited)-1].d + d})

		if u == end {
			lp = max(lp, visited[len(visited)-1].d)
			visited = visited[:len(visited)-1]
		} else {
			stack.PushBack(G[u])
		}

	}
	return lp
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	m := [][]string{}
	i := 0
	for scn.Scan() {
		line := scn.Text()
		gLine := strings.Split(line, "")
		m = append(m, gLine)
		i++
	}

	// part 1
	fmt.Println(bfs(m))

	// part 2
	g := map[pair][]pair{}
	R, C := len(m), len(m[0])

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if m[i][j] != "#" {
				g[pair{i, j}] = getNbrs(i, j, m)
			}
		}
	}
	vertices := []pair{}
	ends := []pair{}
	for k, v := range g {
		if len(v) == 1 || len(v) > 2 {
			vertices = append(vertices, k)
		}

		if len(v) == 1 {
			ends = append(ends, k)
		}
	}

	G := map[pair][]tuple{}
	for _, v := range vertices {
		G[v] = bfs2(g, v)
	}

	start, end := ends[0], ends[1]
	fmt.Println(dfs(G, start, end))
}
