//go:build ignore

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type cube [3]int

func (c cube) isConncted(d cube) bool {
	a := 0
	b := 0
	for i := 0; i < 3; i++ {
		if c[i] == d[i] {
			b++
		} else if c[i] == d[i]+1 || d[i] == c[i]+1 {
			a++
		}
	}
	return a == 1 && b == 2
}

func bfs(cubes map[cube]bool, X, Y, Z int) map[cube]int {
	visited := map[cube]bool{}
	visitedExisting := map[cube]int{}

	q := list.New()
	q.PushBack(cube{0, 0, 0})
	visited[cube{0, 0, 0}] = true

	for q.Len() > 0 {
		curr := q.Front()
		q.Remove(curr)
		c := curr.Value.(cube)

		dx := []int{-1, 1, 0, 0, 0, 0}
		dy := []int{0, 0, -1, 1, 0, 0}
		dz := []int{0, 0, 0, 0, -1, 1}

		for i := 0; i < 6; i++ {
			d := cube{c[0] + dx[i], c[1] + dy[i], c[2] + dz[i]}
			if cubes[d] {
				visitedExisting[d]++
			} else if d[0] < -1 || d[0] > X || d[1] < -1 || d[1] > Y || d[2] < -1 || d[2] > Z || visited[d] {
				continue
			} else {
				visited[d] = true
				q.PushBack(d)
			}
		}
	}

	return visitedExisting
}

func ints(s string) []int {
	p := regexp.MustCompile(`-?\d+`)
	r := []int{}
	for _, e := range p.FindAllString(s, -1) {
		n, _ := strconv.Atoi(e)
		r = append(r, n)
	}
	return r
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	cubes := map[cube]bool{}
	faces := 0
	X, Y, Z := 0, 0, 0
	for scn.Scan() {
		line := scn.Text()
		coords := ints(line)
		x, y, z := coords[0], coords[1], coords[2]
		X, Y, Z = max(X, x), max(Y, y), max(Z, z)

		c := cube{x, y, z}
		curr := 6
		for d := range cubes {
			if c.isConncted(d) {
				curr -= 2
			}
		}
		cubes[c] = true
		faces += curr
	}

	X++
	Y++
	Z++
	fmt.Println(faces)
	fmt.Println(X, Y, Z)
	fmt.Println(bfs(cubes, X, Y, Z))

	r := 0
	for _, k := range bfs(cubes, X, Y, Z) {
		r += k
	}
	fmt.Println(r)
}
