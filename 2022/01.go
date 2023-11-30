package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("inputs/01")

	scanner := bufio.NewScanner(file)

	maxElf := 0
	currElf := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currElf > maxElf {
				maxElf = currElf
			}
			currElf = 0
		} else {
			currCal, _ := strconv.Atoi(line)
			currElf += currCal
		}
	}
	if currElf > maxElf {
		fmt.Println(currElf)
	} else {
		fmt.Println(maxElf)
	}
}
