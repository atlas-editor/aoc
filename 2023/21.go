package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

type tuple struct {
	r, c, d int
}

func isValid(r, c int, garden [][]string, visited map[[2]int]bool) bool {
	rows := len(garden)
	cols := len(garden[0])
	return garden[mod(r, rows)][mod(c, cols)] != "#" && !visited[[2]int{r, c}]
}

func bfs(garden [][]string, start [2]int, it int) int {
	visited := map[[2]int]bool{}
	prev := map[[2]int][2]int{}
	queue := list.New()
	queue.PushBack(tuple{start[0], start[1], 0})
	visited[start] = true

	res := 0
	for queue.Len() > 0 {
		curr := queue.Front()
		queue.Remove(curr)
		currVal := curr.Value.(tuple)
		currPt := [2]int{currVal.r, currVal.c}
		currDis := currVal.d

		if currDis <= it && (currDis+it)%2 == 0 {
			res++
		} else if currDis > it {
			break
		}

		dr := []int{-1, 1, 0, 0}
		dc := []int{0, 0, -1, 1}

		for i := 0; i < 4; i++ {
			newR, newC := currPt[0]+dr[i], currPt[1]+dc[i]
			if isValid(newR, newC, garden, visited) {
				visited[[2]int{newR, newC}] = true
				prev[[2]int{newR, newC}] = currPt
				queue.PushBack(tuple{newR, newC, currDis + 1})
			}
		}
	}
	return res
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	garden := [][]string{}
	start := [2]int{}
	i := 0
	for scn.Scan() {
		line := scn.Text()
		row := strings.Split(line, "")
		garden = append(garden, row)

		if strings.Index(line, "S") != -1 {
			start = [2]int{i, strings.Index(line, "S")}
		}
		i++
	}

	// part 1
	fmt.Println(bfs(garden, start, 64))

	// part 2
	// ...
	R := len(garden)
	steps := 26501365

	m := ((steps - 1) / 2) % R
	f := []int{}
	i = 1
	for len(f) < 3 {
		if ((i-1)/2)%R == m {
			f = append(f, i)
		}
		i += 2
	}

	r0 := bfs(garden, start, f[0])
	r1 := bfs(garden, start, f[1])
	r2 := bfs(garden, start, f[2])

	a := (r2 - r1) - (r1 - r0)
	b := r2 - r1
	N := (steps - f[2]) / (2 * R)

	fmt.Println(r2 + (N * b) + (N * (N + 1) / 2 * a))
}
