package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func init() {
	for k, v := range numericMap {
		numericMapReverse[v] = k
	}

	for k, v := range directionalMap {
		directionalMapReverse[v] = k
	}

	buildNumericPaths()
	buildDirectionalPaths()
}

func main() {
	path := "input.txt"
	//path = "sample.txt"

	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(input string) int {
	s := 0
	for _, keys := range strings.Split(input, "\n") {
		s += solve(keys, 2) * atoi(keys[:len(keys)-1])
	}
	return s
}

func p2(input string) int {
	s := 0
	for _, keys := range strings.Split(input, "\n") {
		s += solve(keys, 25) * atoi(keys[:len(keys)-1])
	}
	return s
}

var numericMap = map[byte]pt{'7': {0, 0}, '8': {0, 1}, '9': {0, 2}, '4': {1, 0}, '5': {1, 1}, '6': {1, 2}, '1': {2, 0}, '2': {2, 1}, '3': {2, 2}, '0': {3, 1}, 'A': {3, 2}}
var numericMapReverse = map[pt]byte{}

var directionalMap = map[byte]pt{'^': {0, 1}, 'A': {0, 2}, '<': {1, 0}, 'v': {1, 1}, '>': {1, 2}}
var directionalMapReverse = map[pt]byte{}

var numericPaths = map[[2]byte][]string{}
var directionalPaths = map[[2]byte][]string{}

func buildNumericPaths() {
	forbidden := pt{3, 0}

	paths := []string{}
	var dfs func(byte, byte, byte, string, set[byte])

	dfs = func(curr, end byte, dir byte, path string, seen set[byte]) {
		if curr == end {
			paths = append(paths, path[1:]+string(dir))
			return
		}

		seen[curr] = true
		path += string(dir)

		currPt := numericMap[curr]
		positions, dirs := nbrs4(currPt[0], currPt[1], 4, 3)
		for i, n := range positions {
			if !seen[numericMapReverse[n]] && n != forbidden {
				dfs(numericMapReverse[n], end, dirs[i], path, seen)
			}
		}

		seen[curr] = false
		path = path[:len(path)-1]
	}

	keys := []byte{'7', '8', '9', '4', '5', '6', '1', '2', '3', '0', 'A'}
	for _, k0 := range keys {
		for _, k1 := range keys {
			if k0 == k1 {
				numericPaths[[2]byte{k0, k1}] = []string{""}
				continue
			}
			paths = []string{}
			dfs(k0, k1, 88, "", set[byte]{})
			//pathsCopy := slices.Clone(paths)

			shortest := 1 << 32
			for _, p := range paths {
				shortest = min(shortest, len(p))
			}

			pathsCopy := []string{}

			for _, p := range paths {
				if len(p) == shortest {
					pathsCopy = append(pathsCopy, p)
				}
			}

			numericPaths[[2]byte{k0, k1}] = pathsCopy
		}
	}
}

func buildDirectionalPaths() {
	forbidden := pt{0, 0}

	paths := []string{}
	var dfs func(byte, byte, byte, string, set[byte])

	dfs = func(curr, end byte, dir byte, path string, seen set[byte]) {
		if curr == end {
			paths = append(paths, path[1:]+string(dir))
			return
		}

		seen[curr] = true
		path += string(dir)

		currPt := directionalMap[curr]
		positions, dirs := nbrs4(currPt[0], currPt[1], 2, 3)
		for i, n := range positions {
			if !seen[directionalMapReverse[n]] && n != forbidden {
				dfs(directionalMapReverse[n], end, dirs[i], path, seen)
			}
		}

		seen[curr] = false
		path = path[:len(path)-1]
	}

	keys := []byte{'^', 'A', '<', 'v', '>'}
	for _, k0 := range keys {
		for _, k1 := range keys {
			if k0 == k1 {
				directionalPaths[[2]byte{k0, k1}] = []string{""}
				continue
			}
			paths = []string{}
			dfs(k0, k1, 88, "", set[byte]{})
			//pathsCopy := slices.Clone(paths)

			shortest := 1 << 32
			for _, p := range paths {
				shortest = min(shortest, len(p))
			}

			pathsCopy := []string{}

			for _, p := range paths {
				if len(p) == shortest {
					pathsCopy = append(pathsCopy, p)
				}
			}

			directionalPaths[[2]byte{k0, k1}] = pathsCopy
		}
	}
}

func buildSeq(input string) []string {
	result := []string{}
	var f func(string, int, byte, string)

	f = func(keys string, index int, prevKey byte, currPath string) {
		if index == len(keys) {
			result = append(result, currPath)
			return
		}

		currKey := keys[index]
		for _, path := range directionalPaths[[2]byte{prevKey, currKey}] {
			f(keys, index+1, keys[index], currPath+path+"A")
		}
	}
	f(input, 0, 'A', "")
	return result
}

func buildSeqNumeric(input string) []string {
	result := []string{}
	var f func(string, int, byte, string)

	f = func(keys string, index int, prevKey byte, currPath string) {
		if index == len(keys) {
			result = append(result, currPath)
			return
		}

		currKey := keys[index]
		for _, path := range numericPaths[[2]byte{prevKey, currKey}] {
			f(keys, index+1, keys[index], currPath+path+"A")
		}
	}
	f(input, 0, 'A', "")
	return result
}

type state struct {
	keys  string
	depth int
}

var cache = map[state]int{}
var re = regexp.MustCompile(`[<>^v]*A`)

func shortestSeq(keys string, depth int) int {
	if depth == 0 {
		return len(keys)
	}
	if v, ok := cache[state{keys, depth}]; ok {
		return v
	}

	subKeys := re.FindAllString(keys, -1)
	total := 0
	for _, subKey := range subKeys {
		subKeySequences := buildSeq(subKey)
		currMin := 1 << 62
		for _, subKeySeq := range subKeySequences {
			currMin = min(currMin, shortestSeq(subKeySeq, depth-1))
		}
		total += currMin
	}

	cache[state{keys, depth}] = total
	return total
}

func solve(input string, d int) int {
	numericSequences := buildSeqNumeric(input)
	min_ := 1 << 62
	for _, seq := range numericSequences {
		x := shortestSeq(seq, d)
		min_ = min(min_, x)
	}
	return min_
}

func nbrs4(r, c, R, C int) ([]pt, []byte) {
	n := []pt{}
	alldirs := []byte{'^', 'v', '<', '>'}
	dirs := []byte{}
	for i, d := range []pt{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		rr, cc := r+d[0], c+d[1]
		if 0 <= rr && rr < R && 0 <= cc && cc < C {
			n = append(n, pt{rr, cc})
			dirs = append(dirs, alldirs[i])
		}
	}
	return n, dirs
}
