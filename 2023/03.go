package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var numRe = regexp.MustCompile(`\d+`)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func containsNotDigitOrDot(s string) bool {
	for _, char := range s {
		if !unicode.IsDigit(char) && char != '.' {
			return true
		}
	}
	return false
}

func isEnginePart(numStart int, numEnd int, prevLine string, currLine string, nextLine string) bool {
	nbhodStart := max(numStart-1, 0)
	nbhoodEnd := min(numEnd+1, len(currLine)-1)
	nbhood := currLine[nbhodStart:nbhoodEnd]

	if prevLine != "" {
		nbhood += prevLine[nbhodStart:nbhoodEnd]
	}
	if nextLine != "" {
		nbhood += nextLine[nbhodStart:nbhoodEnd]
	}

	return containsNotDigitOrDot(nbhood)
}

func gearNbrs(gearIdx int, line string) []int {
	nbrs := []int{}
	for _, m := range numRe.FindAllStringIndex(line, -1) {
		v, _ := strconv.Atoi(line[m[0]:m[1]])
		if gearIdx >= m[0]-1 && gearIdx <= m[1] {
			nbrs = append(nbrs, v)
		}
	}

	return nbrs
}

func processLines(prevLine string, currLine string, nextLine string) (int, int) {
	// part 1
	sum := 0
	for _, m := range numRe.FindAllStringIndex(currLine, -1) {
		match := currLine[m[0]:m[1]]
		v, _ := strconv.Atoi(match)

		if isEnginePart(m[0], m[1], prevLine, currLine, nextLine) {
			sum += v
		}
	}

	// part 2
	sumPart2 := 0
	for i, c := range currLine {
		if c == '*' {
			adjacent := gearNbrs(i, prevLine)
			adjacent = append(adjacent, gearNbrs(i, currLine)...)
			adjacent = append(adjacent, gearNbrs(i, nextLine)...)

			if len(adjacent) == 2 {
				sumPart2 += adjacent[0] * adjacent[1]
			}
		}
	}
	return sum, sumPart2
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	sumPart2 := 0
	var line1, line2, line3 string
	for scanner.Scan() {
		line1, line2, line3 = line2, line3, scanner.Text()

		s1, s2 := processLines(line1, line2, line3)
		sum += s1
		sumPart2 += s2
	}
	s1, s2 := processLines(line2, line3, "")
	sum += s1
	sumPart2 += s2

	fmt.Println(sum, sumPart2)
}
