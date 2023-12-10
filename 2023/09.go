package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isConst(nums []int) (int, bool) {
	r := nums[1] - nums[0]

	for i := 1; i < len(nums)-2; i++ {
		if nums[i+1]-nums[i] != r {
			return -1, false
		}
	}

	return r, true
}

func difs(nums []int) []int {
	res := []int{}

	for i := 0; i < len(nums)-1; i++ {
		res = append(res, nums[i+1]-nums[i])
	}
	return res
}

func findNext(history []int) int {
	q := [][]int{history}

	res := 0
	for {
		curr := q[len(q)-1]
		if v, ok := isConst(curr); ok {
			res = curr[0] - v
			break
		}
		q = append(q, difs(curr))
	}

	for i := len(q) - 2; i >= 0; i-- {
		res = q[i][0] - res
		q[i] = append(q[i], q[i][len(q[i])-1]+res)
	}
	return res

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		strField := strings.Fields(line)
		history := []int{}
		for _, s := range strField {
			v, _ := strconv.Atoi(s)
			history = append(history, v)
		}
		sum += findNext(history)
	}
	fmt.Println(sum)
}
