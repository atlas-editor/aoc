package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type slot struct {
	label string
	fl    int
}

func hash(s string) int {
	res := 0
	for _, r := range s {
		res += int(r)
		res *= 17
		res %= 256

	}
	return res
}

func removeFromMap(ss []slot, key string) []slot {
	idx := -1
	for i, sl := range ss {
		if sl.label == key {
			idx = i
			break
		}
	}

	if idx != -1 {
		return append(ss[:idx], ss[idx+1:]...)
	}
	return ss
}

func addToMap(ss []slot, inpsl slot) []slot {
	found := false
	for i := 0; i < len(ss); i++ {
		if ss[i].label == inpsl.label {
			ss[i].fl = inpsl.fl
			found = true
			break
		}
	}

	if found {
		return ss
	}
	return append(ss, inpsl)
}

func main() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Scan()
	line := scn.Text()
	inps := strings.Split(line, ",")

	m := map[int][]slot{}
	for _, inp := range inps {
		if string(inp[len(inp)-1]) == "-" {
			srem := inp[:len(inp)-1]
			hv := hash(srem)
			m[hv] = removeFromMap(m[hv], srem)
		} else {
			s := strings.Split(inp, "=")
			v, _ := strconv.Atoi(s[1])
			sl := slot{s[0], v}
			shash := hash(s[0])
			m[shash] = addToMap(m[shash], sl)
		}
	}

	ans := 0
	for k, v := range m {
		for j, vv := range v {
			ans += (j + 1) * (k + 1) * vv.fl
		}
	}
	fmt.Println(ans)
}
