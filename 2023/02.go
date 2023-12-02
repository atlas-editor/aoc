package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxCubes := map[string]int{"red": 12, "green": 13, "blue": 14}
	re := regexp.MustCompile(`(\d+)\s(red|green|blue)`)

	i := 0
	sum := 0
	sumPart2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		minCubes := map[string]int{"red": 0, "green": 0, "blue": 0}
		gameOk := true
		for _, m := range re.FindAllStringSubmatch(line, -1) {
			if v, _ := strconv.Atoi(m[1]); maxCubes[m[2]] < v {
				gameOk = false
				// break
			}

			num, _ := strconv.Atoi(m[1])
			color := m[2]
			if minCubes[color] < num {
				minCubes[color] = num
			}

		}
		if gameOk {
			sum += i + 1
		}

		prod := 1
		for _, v := range minCubes {
			prod *= v
		}

		sumPart2 += prod
		i++
	}
	fmt.Println(sum)
	fmt.Println(sumPart2)
}
