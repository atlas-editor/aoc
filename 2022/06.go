package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func counter[S comparable](s []S) map[S]int {
	r := map[S]int{}
	for _, e := range s {
		r[e]++
	}
	return r
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	for i := 0; i < len(input)-14; i++ {
		s := strings.Split(input[i:i+14], "")
		if len(counter(s)) == 14 {
			fmt.Println(i + 14)
			break
		}
	}
}
