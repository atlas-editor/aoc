package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func splitList(l string) []string {
	if l == "[]" {
		return []string{}
	}
	l = l[1 : len(l)-1]
	last := 0
	parts := []string{}
	bs := 0
	for i, e := range l {
		if bs == 0 && e == ',' {
			parts = append(parts, l[last:i])
			last = i + 1
		}
		if e == '[' {
			bs++
		} else if e == ']' {
			bs--
		}
	}
	return append(parts, l[last:len(l)])
}

func compare(left, right string) int {
	intL, errL := strconv.Atoi(left)
	intR, errR := strconv.Atoi(right)

	if errL == nil && errR == nil {
		return intL - intR
	} else if errL == nil {
		return compare("["+left+"]", right)
	} else if errR == nil {
		return compare(left, "["+right+"]")
	}

	partsL := splitList(left)
	partsR := splitList(right)

	for i := 0; i < min(len(partsL), len(partsR)); i++ {
		if r := compare(partsL[i], partsR[i]); r != 0 {
			return r
		}
	}

	return len(partsL) - len(partsR)
}

func main() {
	data, _ := io.ReadAll(os.Stdin)
	pairs := strings.Split(string(data), "\n\n")
	packets := []string{}
	for _, p := range pairs {
		packets = append(packets, strings.Split(p, "\n")...)
	}
	slices.SortFunc(packets, compare)
	fmt.Println(slices.Index(packets, "[[2]]")+1, slices.Index(packets, "[[6]]")+1)
}
