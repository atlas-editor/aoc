package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	digits := map[byte]int{'2': 2, '1': 1, '0': 0, '-': -1, '=': -2}

	sum := 0
	for _, line := range strings.Split(input, "\n") {
		mul := 1
		for i := len(line) - 1; i >= 0; i-- {
			sum += mul * digits[line[i]]
			mul *= 5
		}
	}

	fmt.Println(sum)

	//todo: from decimal to snafu

}
