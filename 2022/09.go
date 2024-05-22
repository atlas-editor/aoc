package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func isClose(head, tail [2]int) bool {
	hy, hx := head[0], head[1]
	ty, tx := tail[0], tail[1]
	if ty >= hy-1 && ty <= hy+1 && tx >= hx-1 && tx <= hx+1 {
		return true
	}
	return false
}

func moveCloser(head, tail [2]int) [2]int {
	hy, hx := head[0], head[1]
	ty, tx := tail[0], tail[1]

	if isClose(head, tail) {
		return tail
	}

	if hy == ty {
		if tx < hx {
			return [2]int{ty, tx + 1}
		}
		return [2]int{ty, tx - 1}
	}

	if hx == tx {
		if ty < hy {
			return [2]int{ty + 1, tx}
		}
		return [2]int{ty - 1, tx}
	}

	if hy < ty {
		if hx < tx {
			return [2]int{ty - 1, tx - 1}
		}
		return [2]int{ty - 1, tx + 1}
	}

	if hx < tx {
		return [2]int{ty + 1, tx - 1}
	}
	return [2]int{ty + 1, tx + 1}
}

var dirMap = map[string][2]int{"R": [2]int{0, 1}, "U": [2]int{1, 0}, "L": [2]int{0, -1}, "D": [2]int{-1, 0}}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	knots := [10][2]int{}
	visited := map[[2]int]bool{knots[9]: true}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		dir, run := fields[0], atoi(fields[1])
		for i := 0; i < run; i++ {
			knots[0] = [2]int{knots[0][0] + dirMap[dir][0], knots[0][1] + dirMap[dir][1]}
			for j := 1; j < 10; j++ {
				knots[j] = moveCloser(knots[j-1], knots[j])
				visited[knots[9]] = true
			}
		}
	}
	fmt.Println(len(visited))
}
