package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func solveQuadratic(a, b, c float64) (float64, float64) {
	D := b*b - 4*a*c
	x0 := (-b + math.Sqrt(D)) / (2 * a)
	x1 := (-b - math.Sqrt(D)) / (2 * a)

	if x0 < x1 {
		return x0, x1
	}

	return x1, x0
}

func intRange(x, y float64) int {
	_, fx := math.Modf(x)
	if fx == 0.0 {
		x++
	}
	_, fy := math.Modf(y)
	if fy == 0.0 {
		y--
	}

	lb := int(math.Ceil(x))
	ub := int(math.Floor(y))

	return ub - lb + 1
}

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()
	scanner := bufio.NewScanner(os.Stdin)

	td := [][]int{}
	for scanner.Scan() {
		numsStr := strings.Fields(scanner.Text())[1:]

		// part 1
		// nums := []int{}
		// for _, n := range numsStr {
		// 	v, _ := strconv.Atoi(n)
		// 	nums = append(nums, v)
		// }
		// td = append(td, nums)

		// part2
		numStr := strings.Join(numsStr, "")
		v, _ := strconv.Atoi(numStr)
		td = append(td, []int{v})
	}

	wins := map[int]int{}
	for i := 0; i < len(td[0]); i++ {
		t := td[0][i]
		d := td[1][i]

		v0, v1 := solveQuadratic(-1.0, float64(t), float64(-d))
		wins[i] = intRange(v0, v1)
	}
	prod := 1
	for _, v := range wins {
		prod *= v
	}
	fmt.Println(prod)
}
