package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type pair struct {
	x, y int
}

func isValid(x, y, rows, cols int, maze [][]string, visited [][]bool) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols && maze[x][y] == "." && !visited[x][y]
}

func bfs(maze [][]string, rows, cols int) [][]bool {
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	queue := list.New()
	// also from input file, point (0,0) is not part of the main loop so it is part of the area outside it
	queue.PushBack(pair{0, 0})
	visited[0][0] = true

	for queue.Len() > 0 {
		curr := queue.Front()
		queue.Remove(curr)
		currPt := curr.Value.(pair)

		dx := []int{-1, 1, 0, 0}
		dy := []int{0, 0, -1, 1}

		for i := 0; i < 4; i++ {
			newX, newY := currPt.x+dx[i], currPt.y+dy[i]
			if isValid(newX, newY, rows, cols, maze, visited) {
				visited[newX][newY] = true
				queue.PushBack(pair{newX, newY})
			}
		}
	}

	return visited
}

func main() {
	pipeDirs := map[rune]map[pair]pair{}
	pipeDirs['|'] = map[pair]pair{pair{1, 0}: pair{1, 0}}
	pipeDirs['|'][pair{-1, 0}] = pair{-1, 0}
	pipeDirs['-'] = map[pair]pair{pair{0, 1}: pair{0, 1}}
	pipeDirs['-'][pair{0, -1}] = pair{0, -1}
	pipeDirs['L'] = map[pair]pair{pair{1, 0}: pair{0, 1}}
	pipeDirs['L'][pair{0, -1}] = pair{-1, 0}
	pipeDirs['J'] = map[pair]pair{pair{1, 0}: pair{0, -1}}
	pipeDirs['J'][pair{0, 1}] = pair{-1, 0}
	pipeDirs['7'] = map[pair]pair{pair{0, 1}: pair{1, 0}}
	pipeDirs['7'][pair{-1, 0}] = pair{0, -1}
	pipeDirs['F'] = map[pair]pair{pair{0, -1}: pair{1, 0}}
	pipeDirs['F'][pair{-1, 0}] = pair{0, 1}

	scanner := bufio.NewScanner(os.Stdin)

	maze := map[pair]rune{}
	i := 0
	lineLen := 0
	var start pair
	for scanner.Scan() {
		line := scanner.Text()
		lineLen = len(line)
		for j := 0; j < len(line); j++ {
			maze[pair{i, j}] = rune(line[j])
		}
		idx := strings.IndexRune(line, 'S')
		if idx > -1 {
			start = pair{i, idx}
		}
		i++
	}
	rows, cols := lineLen, i

	isOnMainLoop := map[pair]bool{start: true}
	// hardcoded from input ¯\_(ツ)_/¯ find 'S' and pick a direction and the corresponding tile to adjust
	pos := pair{start.x + 1, start.y}
	dir := pair{0, 1}
	pipe := 'L'
	length := 1
	for pipe != 'S' {
		isOnMainLoop[pos] = true
		length++
		pos.x += dir.x
		pos.y += dir.y
		pipe = maze[pos]
		dir = pipeDirs[pipe][dir]
	}
	maze[start] = '|'

	extendedMaze := make([][]string, 2*cols)
	for i := range extendedMaze {
		extendedMaze[i] = make([]string, 2*rows)
	}

	mazeJustLoop := make([][]string, cols)
	for i := range mazeJustLoop {
		mazeJustLoop[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		var currPos pair
		for j := 0; j < cols; j++ {
			currPos = pair{i, j}
			if _, found := isOnMainLoop[currPos]; found {
				extendedMaze[2*i][2*j] = "#"
				mazeJustLoop[i][j] = "#"
				if maze[currPos] == '-' || maze[currPos] == 'L' || maze[currPos] == 'F' {
					extendedMaze[2*i][2*j+1] = "#"
				} else {
					extendedMaze[2*i][2*j+1] = "."
				}
			} else {
				extendedMaze[2*i][2*j] = "."
				extendedMaze[2*i][2*j+1] = "."
				mazeJustLoop[i][j] = "."
			}
		}

		for j := 0; j < cols; j++ {
			currPos = pair{i, j}
			if _, found := isOnMainLoop[currPos]; found {
				if maze[currPos] == '|' || maze[currPos] == 'F' || maze[currPos] == '7' {
					extendedMaze[2*i+1][2*j] = "#"
				} else {
					extendedMaze[2*i+1][2*j] = "."
				}
			} else {
				extendedMaze[2*i+1][2*j] = "."
			}
			extendedMaze[2*i+1][2*j+1] = "."
		}
	}

	outsideMainLoop := bfs(extendedMaze, 2*rows, 2*cols)

	sum := 0
	for i, l := range mazeJustLoop {
		for j, _ := range l {
			if !isOnMainLoop[pair{i, j}] && !outsideMainLoop[2*i][2*j] {
				sum++
			}
		}
	}

	fmt.Println(sum)
}
