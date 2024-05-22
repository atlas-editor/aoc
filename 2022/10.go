package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cycle := 1
	X := 1
	m := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		// if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		// 	fmt.Println(X * cycle)
		// }
		m[cycle] = X
		if strings.HasPrefix(line, "noop") {
			cycle++
			continue
		}

		f := strings.Fields(line)
		v, _ := strconv.Atoi(f[1])

		m[cycle+1] = X
		X += v
		cycle += 2
	}

	sum := 0
	for row := 0; row < 6; row++ {
		for j := 1; j <= 40; j++ {
			cycle := (40 * row) + j
			spritePos := m[cycle]
			// fmt.Println(j, spritePos)
			if spritePos > j-3 && spritePos < j+1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println(m)
	fmt.Println(sum)
}
