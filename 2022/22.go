//go:build ignore

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var dirCost = map[byte]int{'E': 0, 'S': 1, 'W': 2, 'N': 3}

func main() {
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	input := string(data)
	parts := strings.Split(input, "\n\n")

	map_ := strings.Split(parts[0], "\n")
	directions := parts[1]
	pathRe := regexp.MustCompile(`(\d+|[LR])`)

	currR := 0
	currC := 50
	currDir := byte('E')
	for _, m := range pathRe.FindAllStringSubmatch(directions, -1) {
		if steps, err := strconv.Atoi(m[1]); err == nil {
			for range steps {
				posR, posC, d := move(currR, currC, currDir)
				if map_[posR][posC] == '#' {
					break
				}

				currR = posR
				currC = posC
				currDir = d
			}
		} else {
			currDir = turn(currDir, m[1][0])
		}
	}

	fmt.Println(1000*(currR+1) + 4*(currC+1) + dirCost[currDir])
}

func turn(dir byte, orient byte) byte {
	switch dir {
	case 'W':
		if orient == 'L' {
			return 'S'
		} else {
			return 'N'
		}
	case 'S':
		if orient == 'L' {
			return 'E'
		} else {
			return 'W'
		}
	case 'E':
		if orient == 'L' {
			return 'N'
		} else {
			return 'S'
		}
	case 'N':
		if orient == 'L' {
			return 'W'
		} else {
			return 'E'
		}
	default:
		panic("invalid dir")
	}
}

func move(r, c int, dir byte) (int, int, byte) {
	// part 2

	// 1 (0-49, 50)
	if 0 <= r && r <= 49 && c == 50 && dir == 'W' {
		return 149 - r, 0, 'E'
	}
	// 4 (100-149, 0)
	if 100 <= r && r <= 149 && c == 0 && dir == 'W' {
		return 149 - r, 50, 'E'
	}

	// 2 (50-99, 50)
	if 50 <= r && r <= 99 && c == 50 && dir == 'W' {
		return 100, r - 50, 'S'
	}
	// 3 (100, 0-49)
	if r == 100 && 0 <= c && c <= 49 && dir == 'N' {
		return c + 50, 50, 'E'
	}

	// 5 (150-199, 0)
	if 150 <= r && r <= 199 && c == 0 && dir == 'W' {
		return 0, r - 100, 'S'
	}
	// 14 (0, 50-99)
	if r == 0 && 50 <= c && c <= 99 && dir == 'N' {
		return c + 100, 0, 'E'
	}

	// 6 (199, 0-49)
	if r == 199 && 0 <= c && c <= 49 && dir == 'S' {
		return 0, c + 100, 'S'
	}
	// 13 (0, 100-149)
	if r == 0 && 100 <= c && c <= 149 && dir == 'N' {
		return 199, c - 100, 'N'
	}

	// 7 (150-199, 49)
	if 150 <= r && r <= 199 && c == 49 && dir == 'E' {
		return 149, r - 100, 'N'
	}
	// 8 (149, 50-99)
	if r == 149 && 50 <= c && c <= 99 && dir == 'S' {
		return c + 100, 49, 'W'
	}

	// 9 (100-149, 99)
	if 100 <= r && r <= 149 && c == 99 && dir == 'E' {
		return 149 - r, 149, 'W'
	}
	// 12 (0-49, 149)
	if 0 <= r && r <= 49 && c == 149 && dir == 'E' {
		return 149 - r, 99, 'W'
	}

	// 10 (50-99, 99)
	if 50 <= r && r <= 99 && c == 99 && dir == 'E' {
		return 49, r + 50, 'N'
	}
	// 11 (49, 100-149)
	if r == 49 && 100 <= c && c <= 149 && dir == 'S' {
		return c - 50, 99, 'W'
	}

	// part1
	//	if 0 <= r && r <= 49 && c == 50 && dir == 'W' {
	//		return r, 149
	//	}
	//	if 50 <= r && r <= 99 && c == 50 && dir == 'W' {
	//		return r, 99
	//	}
	//	if 100 <= r && r <= 149 && c == 0 && dir == 'W' {
	//		return r, 99
	//	}
	//	if 150 <= r && r <= 199 && c == 0 && dir == 'W' {
	//		return r, 49
	//	}
	//
	//
	//	if r == 199 && 0 <= c && c <= 49 && dir == 'S' {
	//		return 100, c
	//	}
	//	if r == 149 && 50 <= c && c <= 99 && dir == 'S' {
	//		return 0, c
	//	}
	//	if r == 49 && 100 <= c && c <= 149 && dir == 'S' {
	//		return 0, c
	//	}
	//
	//
	//	if 150 <= r && r <= 199 && c == 49 && dir == 'E' {
	//		return r, 0
	//	}
	//	if 100 <= r && r <= 149 && c == 99 && dir == 'E' {
	//		return r, 0
	//	}
	//	if 50 <= r && r <= 99 && c == 99 && dir == 'E' {
	//		return r, 50
	//	}
	//	if 0 <= r && r <= 49 && c == 149 && dir == 'E' {
	//		return r, 50
	//	}
	//
	//	if r == 0 && 100 <= c && c <= 149 && dir == 'N' {
	//		return 49, c
	//	}
	//	if r == 0 && 50 <= c && c <= 99 && dir == 'N' {
	//		return 149, c
	//	}
	//	if r == 100 && 0 <= c && c <= 49 && dir == 'N' {
	//		return 199, c
	//	}

	switch dir {
	case 'W':
		return r, c - 1, 'W'
	case 'S':
		return r + 1, c, 'S'
	case 'E':
		return r, c + 1, 'E'
	case 'N':
		return r - 1, c, 'N'
	default:
		panic("invalid dir")
	}
}
