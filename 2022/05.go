package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func transpose[S any](m [][]S) [][]S {
	rows := len(m)
	cols := len(m[0])

	r := make([][]S, cols)
	for i := 0; i < cols; i++ {
		r[i] = make([]S, rows)
		for j := 0; j < rows; j++ {
			r[i][j] = m[j][i]
		}
	}
	return r
}

func slicesMap[S, T any](ts []S, f func(S) T) []T {
	us := []T{}
	for _, e := range ts {
		us = append(us, f(e))
	}
	return us
}

func moveCrate(crates [][]string, i, j int) [][]string {
	crates[j] = append(crates[j], crates[i][len(crates[i])-1])
	crates[i] = crates[i][:len(crates[i])-1]
	return crates
}

func moveCrates(crates [][]string, no, i, j int) [][]string {
	c := crates

	crates[j] = append(crates[j], crates[i][len(crates[i])-no:]...)
	crates[i] = crates[i][:len(crates[i])-no]
	return c
}

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	input, _ := io.ReadAll(os.Stdin)
	parts := strings.Split(string(input), "\n\n")
	crates := strings.Split(parts[0], "\n")
	crates = crates[:len(crates)-1]
	cratesElems := slicesMap(crates, func(s string) []string { return strings.Split(s, "") })
	cratesElems = transpose(cratesElems)
	cratesElems = cratesElems[1 : len(cratesElems)-1]
	columns := [][]string{}
	for i := 0; i < len(cratesElems); i += 4 {
		columns = append(columns, cratesElems[i])
	}

	stacks := [][]string{}
	stacks = append(stacks, []string{})
	for idx, c := range columns {
		stacks = append(stacks, []string{})
		for i := len(c) - 1; i >= 0; i-- {
			if strings.TrimSpace(c[i]) == "" {
				break
			}
			stacks[idx+1] = append(stacks[idx+1], c[i])
		}
	}
	for _, l := range strings.Split(parts[1], "\n") {
		data := strings.Fields(l)
		no, i, j := atoi(data[1]), atoi(data[3]), atoi(data[5])
		stacks = moveCrates(stacks, no, i, j)
	}

	for i := 1; i < len(stacks); i++ {
		fmt.Print(stacks[i][len(stacks[i])-1])
	}
	fmt.Println()
}
