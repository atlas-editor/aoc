package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "input.txt"
	//path = "sample.txt"

	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(input string) int {
	return -1
}

func p2(input string) int {
	return -1
}
