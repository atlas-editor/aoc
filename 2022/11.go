package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func findTwoLargest(numbers []int) (int, int) {
	slices.Sort(numbers)
	slices.Reverse(numbers)
	return numbers[0], numbers[1]
}

type Monkey struct {
	name  int
	items []int
	op    func(int) int
	test  func(int) bool
	next  func(bool) int
}

func main() {
	data, _ := io.ReadAll(os.Stdin)
	input := string(data)

	x := strings.Split(input, "\n\n")

	ms := map[int]Monkey{}
	lcm := 1
	for i, m := range x {
		info := strings.Split(m, "\n")

		_, itemsLine, _ := strings.Cut(info[1], ": ")
		itemsField := strings.Split(itemsLine, ", ")
		items := []int{}
		for _, it := range itemsField {
			v, _ := strconv.Atoi(it)
			items = append(items, v)
		}

		_, opsLine, _ := strings.Cut(info[2], "old ")
		o := opsLine[0]
		num := opsLine[2:]
		var op func(int) int
		if num == "old" {
			op = func(x int) int { return x * x }
		} else {
			v, _ := strconv.Atoi(num)
			if rune(o) == '+' {
				op = func(x int) int { return x + v }
			} else {
				op = func(x int) int { return x * v }
			}
		}

		_, testLine, _ := strings.Cut(info[3], "by ")
		v, _ := strconv.Atoi(testLine)
		lcm *= v
		test := func(x int) bool { return x%v == 0 }
		_, ifTrueStr, _ := strings.Cut(info[4], "monkey ")
		_, ifFalseStr, _ := strings.Cut(info[5], "monkey ")
		z, _ := strconv.Atoi(ifTrueStr)
		zz, _ := strconv.Atoi(ifFalseStr)

		next := func(x bool) int {
			if x {
				return z
			} else {
				return zz
			}
		}

		monkey := Monkey{i, items, op, test, next}
		ms[i] = monkey

	}
	sum := map[int]int{}
	for k := 0; k < 10000; k++ {
		for i := 0; i < len(ms); i++ {
			m := ms[i]
			// fmt.Println(k, i, len(m.items))
			for _, it := range m.items {
				worry := m.op(it) % lcm
				n := m.next(m.test(worry))
				// fmt.Println(k, i, it, worry, m.test(worry), n)
				nM := ms[n]
				nM.items = append(nM.items, worry)
				ms[n] = nM
				sum[i]++
			}
			m.items = []int{}
			ms[i] = m
			// fmt.Println()
		}
	}
	s := []int{}
	for _, v := range sum {
		s = append(s, v)
	}

	m0, m1 := findTwoLargest(s)
	for k, v := range sum {
		fmt.Println(k, v)
	}
	fmt.Println(m0 * m1)
}
