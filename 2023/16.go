package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type tuple struct {
	x, y int
	dir  [2]int
}

func isInside(x, y, rows, cols int, maze [][]string) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func next(x, y, rows, cols int, dir [2]int, maze [][]string) []tuple {
	nexts := []tuple{}
	if maze[y][x] == "." {
		if isInside(x-dir[0], y-dir[1], rows, cols, maze) {
			return []tuple{tuple{x - dir[0], y - dir[1], dir}}
		} else {
			return nexts
		}
	}

	if maze[y][x] == "|" {
		if dir == [2]int{0, -1} || dir == [2]int{0, 1} {
			if isInside(x-dir[0], y-dir[1], rows, cols, maze) {
				return []tuple{tuple{x - dir[0], y - dir[1], dir}}
			} else {
				return nexts
			}
		} else {
			dir1, dir2 := [2]int{0, -1}, [2]int{0, 1}
			if isInside(x-dir1[0], y-dir1[1], rows, cols, maze) {
				nexts = append(nexts, tuple{x - dir1[0], y - dir1[1], dir1})
			}
			if isInside(x-dir2[0], y-dir2[1], rows, cols, maze) {
				nexts = append(nexts, tuple{x - dir2[0], y - dir2[1], dir2})
			}
			return nexts
		}
	}

	if maze[y][x] == "-" {
		if dir == [2]int{-1, 0} || dir == [2]int{1, 0} {
			if isInside(x-dir[0], y-dir[1], rows, cols, maze) {
				return []tuple{tuple{x - dir[0], y - dir[1], dir}}
			} else {
				return nexts
			}
		} else {
			dir1, dir2 := [2]int{-1, 0}, [2]int{1, 0}
			if isInside(x-dir1[0], y-dir1[1], rows, cols, maze) {
				nexts = append(nexts, tuple{x - dir1[0], y - dir1[1], dir1})
			}
			if isInside(x-dir2[0], y-dir2[1], rows, cols, maze) {
				nexts = append(nexts, tuple{x - dir2[0], y - dir2[1], dir2})
			}
			return nexts
		}
	}

	nextDir := [2]int{}
	if maze[y][x] == "\\" {
		if dir == [2]int{-1, 0} {
			nextDir = [2]int{0, -1}
		}
		if dir == [2]int{1, 0} {
			nextDir = [2]int{0, 1}
		}
		if dir == [2]int{0, -1} {
			nextDir = [2]int{-1, 0}
		}
		if dir == [2]int{0, 1} {
			nextDir = [2]int{1, 0}
		}

		if isInside(x-nextDir[0], y-nextDir[1], rows, cols, maze) {
			return []tuple{tuple{x - nextDir[0], y - nextDir[1], nextDir}}
		}
	}

	if maze[y][x] == "/" {
		if dir == [2]int{-1, 0} {
			nextDir = [2]int{0, 1}
		}
		if dir == [2]int{1, 0} {
			nextDir = [2]int{0, -1}
		}
		if dir == [2]int{0, -1} {
			nextDir = [2]int{1, 0}
		}
		if dir == [2]int{0, 1} {
			nextDir = [2]int{-1, 0}
		}

		if isInside(x-nextDir[0], y-nextDir[1], rows, cols, maze) {
			return []tuple{tuple{x - nextDir[0], y - nextDir[1], nextDir}}
		}
	}

	return []tuple{}
}

func bfs(maze [][]string, start tuple) [][]bool {
	rows := len(maze)
	cols := len(maze[0])
	visited := map[tuple]bool{}

	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	for queue.Len() > 0 {
		curr := queue.Front()
		queue.Remove(curr)
		currPos := curr.Value.(tuple)

		for _, nextPos := range next(currPos.x, currPos.y, rows, cols, currPos.dir, maze) {
			if visited[nextPos] {
				continue
			}
			visited[nextPos] = true
			queue.PushBack(nextPos)
		}
	}

	visitedPts := make([][]bool, rows)
	for i := range visitedPts {
		visitedPts[i] = make([]bool, cols)
	}
	for k, _ := range visited {
		visitedPts[k.y][k.x] = true
	}
	return visitedPts
}

func printVisited(r [][]bool) {
	for _, l := range r {
		for _, ch := range l {
			if ch == true {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func countTrue(r [][]bool) int {
	res := 0
	for _, l := range r {
		for _, ch := range l {
			if ch == true {
				res++
			}
		}
	}
	return res
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	m := [][]string{}
	for scn.Scan() {
		line := scn.Text()
		chars := strings.Split(line, "")
		m = append(m, chars)
	}

	res := 0
	w := len(m)
	for i := 0; i < 4; i++ {
		for j := 0; j < w; j++ {
			start := tuple{}
			if i == 0 {
				start = tuple{j, 0, [2]int{0, -1}}
			} else if i == 1 {
				start = tuple{w - 1, j, [2]int{1, 0}}
			} else if i == 2 {
				start = tuple{j, w - 1, [2]int{0, 1}}
			} else if i == 3 {
				start = tuple{0, j, [2]int{-1, 0}}
			}
			curr := countTrue(bfs(m, start))
			res = max(res, curr)
		}
	}
	fmt.Println(res)
}
