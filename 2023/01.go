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

func createNumMap() (map[string]string, map[string]string) {
	numMapTmp := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	numMap := map[string]string{}
	reverseKeysMap := map[string]string{}
	i := 1
	for k, v := range numMapTmp {
		numMap[k] = v
		numMap[fmt.Sprint(i)] = fmt.Sprint(i)
		reverseKeysMap[k] = reverseString(k)
		reverseKeysMap[fmt.Sprint(i)] = fmt.Sprint(i)
		i += 1
	}

	return numMap, reverseKeysMap
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numMap, reverseKeysMap := createNumMap()
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		reversedLine := reverseString(line)
		lineLen := len(line)

		first := ""
		last := ""
	LineLoop:
		for i := 0; i < lineLen; i++ {
			for k, v := range numMap {
				if first != "" && last != "" {
					break LineLoop
				}

				l := len(k)
				if first == "" && i+l <= lineLen && line[i:i+l] == k {
					first = v
				}
				if last == "" && i+l <= lineLen && reversedLine[i:i+l] == reverseKeysMap[k] {
					last = v
				}
			}
		}
		v, _ := strconv.Atoi(first + last)
		sum += v
	}
	fmt.Println(sum)
}
