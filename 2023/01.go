package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func createNumMap() map[string]string {
	numMap := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	for i := 1; i < 10; i++ {
		iStr := strconv.Itoa(i)
		numMap[iStr] = iStr
	}
	return numMap
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numMap := createNumMap()
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		first := ""
		last := ""
	LineLoop:
		for i := 0; i < lineLen; i++ {
			j := lineLen - i
			for k, v := range numMap {
				if first != "" && last != "" {
					break LineLoop
				}

				l := len(k)
				if first == "" && i+l <= lineLen && line[i:i+l] == k {
					first = v
				}
				if last == "" && j-l >= 0 && line[j-l:j] == k {
					last = v
				}
			}
		}
		v, _ := strconv.Atoi(first + last)
		sum += v
	}
	fmt.Println(sum)
}
