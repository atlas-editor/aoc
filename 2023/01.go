package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func reverseString(s string) string {
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	return reversed
}

func createNumMap() map[string]string {
	numMapTmp := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	numMap := map[string]string{}
	i := 1
	for k, v := range numMapTmp {
		numMap[k] = v
		numMap[reverseString(k)] = v
		numMap[fmt.Sprint(i)] = fmt.Sprint(i)
		i += 1
	}

	return numMap
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numMap := createNumMap()

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		first := ""
		last := ""

		for i := 0; i < len(line); i++ {
			for k, v := range numMap {
				l := len(k)
				if i+l <= len(line) && line[i:i+l] == k {
					first = v
					break
				}
			}
			if first != "" {
				break
			}
		}

		reversedLine := reverseString(line)
		for i := 0; i < len(reversedLine); i++ {
			for k, v := range numMap {
				l := len(k)
				if i+l <= len(reversedLine) && reversedLine[i:i+l] == k {
					last = v
					break
				}
			}
			if last != "" {
				break
			}
		}
		v, _ := strconv.Atoi(first + last)
		sum += v
	}

	fmt.Println(sum)
}
