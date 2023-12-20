package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func makeKey(s string, b []int) string {
	bIntSlice := []string{}
	for _, v := range b {
		bIntSlice = append(bIntSlice, strconv.Itoa(v))
	}

	return s + " " + strings.Join(bIntSlice, ",")
}

func noDot(s string, v int) (bool, string) {
	if v > len(s) {
		return false, ""
	}

	for i := 0; i < v; i++ {
		if rune(s[i]) == '.' {
			return false, ""
		}
	}

	if len(s) > v {
		if rune(s[v]) == '#' {
			return false, ""
		} else {
			return true, s[v+1:]
		}
	}

	return true, ""
}

func dp(s string, g []int) int {
	if v, found := cache[makeKey(s, g)]; found {
		return v
	}

	if len(g) == 0 {
		if strings.IndexRune(s, '#') == -1 {
			cache[makeKey(s, g)] = 1
			return 1
		}
		cache[makeKey(s, g)] = 0
		return 0
	}

	if len(s) == 0 {
		cache[makeKey(s, g)] = 0
		return 0
	}
	r := rune(s[0])

	if r == '.' {
		res := dp(s[1:], g)
		cache[makeKey(s, g)] = res
		return res
	}

	if r == '#' {
		if ok, ss := noDot(s, g[0]); ok {
			res := dp(ss, g[1:])
			cache[makeKey(s, g)] = res
			return res
		}
		cache[makeKey(s, g)] = 0
		return 0
	}

	if r == '?' {
		res := dp("."+s[1:], g) + dp("#"+s[1:], g)
		cache[makeKey(s, g)] = res
		return res
	}

	return 0

}

func duplStr(s string, rep int, sep string) string {
	resSlice := []string{}

	for i := 0; i < rep; i++ {
		resSlice = append(resSlice, s)
	}

	return strings.Join(resSlice, sep)
}

func duplInts(nums []int, rep int) (res []int) {
	for i := 0; i < rep; i++ {
		for _, n := range nums {
			res = append(res, n)
		}
	}
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	res := 0
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		r := line[0]
		rrStr := strings.Split(line[1], ",")
		rr := []int{}
		for _, ru := range rrStr {
			v, _ := strconv.Atoi(string(ru))
			rr = append(rr, v)
		}

		r2 := duplStr(r, 5, "?")
		rr2 := duplInts(rr, 5)

		res += dp(r2, rr2)
	}
	fmt.Println(res)
}
