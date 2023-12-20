package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func transpose(input [][]string) [][]string {
	if len(input) == 0 || len(input[0]) == 0 {
		return nil
	}

	rows := len(input[0])
	cols := len(input)

	transposed := make([][]string, rows)
	for i := range transposed {
		transposed[i] = make([]string, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			transposed[i][j] = input[j][i]
		}
	}

	return transposed
}

// for part1 replace isAlmostRefl by this func
func isRefl(isl [][]string, idx int) bool {
	a := idx - 1
	b := idx + 2

	for a >= 0 && b <= len(isl)-1 {
		if !slices.Equal(isl[a], isl[b]) {
			return false
		}
		a--
		b++
	}
	return true
}

func slicesDiff(s0 []string, s1 []string) int {
	res := 0
	for i := 0; i < len(s0); i++ {
		if s0[i] != s1[i] {
			res++
		}
	}
	return res
}

func isAlmostRefl(isl [][]string, idx int) bool {
	a := idx
	b := idx + 1
	diff := 0

	for a >= 0 && b <= len(isl)-1 && diff <= 1 {
		diff += slicesDiff(isl[a], isl[b])
		a--
		b++
	}
	if diff != 1 {
		return false
	}
	return true
}

func processIsland(isl [][]string) int {
	res := 0
	for i := 0; i < len(isl)-1; i++ {
		if isAlmostRefl(isl, i) {
			res += i + 1
		}
	}
	return res
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	island := [][]string{}
	res := 0
	for scn.Scan() {
		line := scn.Text()
		if len(strings.TrimSpace(line)) == 0 {
			res += processIsland(island)*100 + processIsland(transpose(island))
			island = [][]string{}
			continue
		}
		sl := strings.Split(line, "")
		island = append(island, sl)
	}

	if len(island) > 0 {
		res += processIsland(island)*100 + processIsland(transpose(island))
	}

	fmt.Println(res)
}
