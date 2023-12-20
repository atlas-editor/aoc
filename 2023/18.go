package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func areaPolygon(corners []pair) int {
	a := 0
	for i := 0; i < len(corners); i++ {
		p := corners[i]
		q := corners[(i+1)%len(corners)]
		a += -p.r*q.c + p.c*q.r
	}
	return a / 2
}

type pair struct {
	r, c int
}

func main() {
	var U pair = pair{-1, 0}
	var L pair = pair{0, -1}
	var D pair = pair{1, 0}
	var R pair = pair{0, 1}

	dirMap := map[string]pair{"U": U, "L": L, "D": D, "R": R}

	scn := bufio.NewScanner(os.Stdin)

	dirs := []string{}
	lens := []int{}
	part1 := false
	for scn.Scan() {
		line := scn.Text()
		d := strings.Fields(line)

		dirStr, nStr, hexStr := d[0], d[1], d[2]
		n := 0

		if part1 {
			n, _ = strconv.Atoi(nStr)
		} else {
			hexStr = hexStr[2 : len(hexStr)-1]
			nStr = hexStr[:len(hexStr)-1]
			n64, _ := strconv.ParseInt(nStr, 16, 64)
			n = int(n64)

			dirStr = hexStr[len(hexStr)-1:]

			if dirStr == "0" {
				dirStr = "R"
			} else if dirStr == "1" {
				dirStr = "D"
			} else if dirStr == "2" {
				dirStr = "L"
			} else if dirStr == "3" {
				dirStr = "U"
			}
		}

		dirs = append(dirs, dirStr)
		lens = append(lens, n)
	}

	N := len(dirs)

	last := pair{0, 0}
	corners := []pair{}
	for i := 0; i < N; i++ {
		dir, n := dirs[i], lens[i]
		nextDir := dirs[(i+1)%N]

		dirPair := dirMap[dir]

		next := pair{last.r + (dirPair.r * n), last.c + (dirPair.c * n)}

		cornerType := dir + nextDir
		currCorner := pair{}
		// we assume the enumeration of turning pts is clockwise
		if cornerType == "RD" {
			currCorner = pair{next.r, next.c + 1}
		}
		if cornerType == "RU" {
			currCorner = pair{next.r, next.c}
		}
		if cornerType == "LD" {
			currCorner = pair{next.r + 1, next.c + 1}
		}
		if cornerType == "LU" {
			currCorner = pair{next.r + 1, next.c}
		}
		if cornerType == "DR" {
			currCorner = pair{next.r, next.c + 1}
		}
		if cornerType == "UR" {
			currCorner = pair{next.r, next.c}
		}
		if cornerType == "DL" {
			currCorner = pair{next.r + 1, next.c + 1}
		}
		if cornerType == "UL" {
			currCorner = pair{next.r + 1, next.c}
		}

		corners = append(corners, currCorner)
		last = next
	}
	fmt.Println(areaPolygon(corners))
}
