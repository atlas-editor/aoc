package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]

	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

const (
	X = 101
	Y = 103
)

func p1(input string) int {
	q := map[int]int{}
	for _, line := range strings.Split(input, "\n") {
		nums := ints(line)
		px, py, vx, vy := nums[0], nums[1], nums[2], nums[3]
		r := robot{vec{px, py}, vec{vx, vy}}

		pos := simulate(r, 100)
		q[quadrant(pos)]++
	}

	return q[0] * q[1] * q[2] * q[3]
}

func p2(input string) int {
	robots := []robot{}
	for _, line := range strings.Split(input, "\n") {
		nums := ints(line)
		px, py, vx, vy := nums[0], nums[1], nums[2], nums[3]
		r := robot{vec{px, py}, vec{vx, vy}}

		robots = append(robots, r)
	}

	i := 0
	for {
		if hasTree(robots) {
			return i
		}
		for j := range robots {
			robots[j].p = simulate(robots[j], 1)
		}
		i++
	}
}

type robot struct {
	p, v vec
}

func simulate(r robot, steps int) vec {
	pos := r.p
	for range steps {
		pos = pos.add(r.v)
	}
	return vec{((pos[0] % X) + X) % X, ((pos[1] % Y) + Y) % Y}
}

func quadrant(pos vec) int {
	midX, midY := X/2, Y/2
	x, y := pos[0], pos[1]
	switch {
	case x < midX && y < midY:
		return 0
	case x > midX && y < midY:
		return 1
	case x < midX && y > midY:
		return 2
	case x > midX && y > midY:
		return 3
	default:
		return -1
	}
}

func hasTree(robots []robot) bool {
	robotMap := [Y][X]bool{}

	for _, r := range robots {
		x, y := r.p[0], r.p[1]
		robotMap[y][x] = true
	}

	seen := set[vec]{}
	for y := range Y {
		for x := range X {
			if seen[vec{x, y}] && !robotMap[y][x] {
				continue
			}
			seen[vec{x, y}] = true
			size := 0
			q := []vec{{x, y}}
			for len(q) > 0 {
				curr := pop(&q)
				size++
				for _, n := range nbrs4(curr[0], curr[1]) {
					if !seen[n] && robotMap[n[1]][n[0]] {
						q = append(q, n)
						seen[n] = true
					}
				}
			}
			if size > 100 {
				return true
			}
		}
	}

	return false
}

func nbrs4(x, y int) []vec {
	nbrs := []vec{}
	for _, d := range []vec{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		xx, yy := x+d[0], y+d[1]
		if inBound(xx, yy) {
			nbrs = append(nbrs, vec{xx, yy})
		}
	}
	return nbrs
}

func inBound(x, y int) bool {
	return 0 <= x && x < X && 0 <= y && y < Y
}

/*
utils
*/

type vec [2]int

func (u vec) add(v vec) vec {
	return vec{u[0] + v[0], u[1] + v[1]}
}

type set[T comparable] map[T]bool

func pop[T any](slice *[]T) T {
	n := len(*slice)
	if n == 0 {
		panic("empty slice")
	}
	back := (*slice)[n-1]
	*slice = (*slice)[:n-1]
	return back
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func ints(s string) (r []int) {
	p := regexp.MustCompile(`-?\d+`)
	for _, e := range p.FindAllString(s, -1) {
		r = append(r, atoi(e))
	}
	return
}
