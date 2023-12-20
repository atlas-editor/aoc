package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type part [4]int // x, m, a, s ..ints

type intervals [4][2]int

type mapPos struct {
	name string
	idx  int
}

// for part 1 only
func wfFunc(p part, wf string) (string, bool) {
	condTarget := strings.Split(wf, ":")
	if len(condTarget) == 1 {
		return condTarget[0], true
	}
	cond := condTarget[0]
	target := condTarget[1]

	partType := string(cond[0])
	idx := strings.Index("xmas", partType)
	op := string(cond[1])
	val, _ := strconv.Atoi(cond[2:])

	if op == "<" {
		return target, p[idx] < val
	}
	return target, p[idx] > val
}

func getTarget(wf string) string {
	t := strings.Split(wf, ":")
	if len(t) == 1 {
		return wf
	}
	return t[1]
}

func process(acc mapPos, wfs []string, r intervals) intervals {
	first := wfs[acc.idx]
	if strings.Index(first, ":") != -1 {
		tmp := strings.Split(first, ":")
		cond := tmp[0]
		r = reduceIntervals(r, cond, false)
	}
	for i := acc.idx - 1; i >= 0; i-- {
		curr := wfs[i]
		tmp := strings.Split(curr, ":")
		cond := tmp[0]
		r = reduceIntervals(r, cond, true)
	}
	return r
}

func overlapInterval(i0, i1 [2]int) [2]int {
	return [2]int{max(i0[0], i1[0]), min(i0[1], i1[1])}
}

func reduceIntervals(r intervals, cond string, reverse bool) intervals {
	idx := strings.Index("xmas", cond[:1])
	op := cond[1:2]
	val, _ := strconv.Atoi(cond[2:])
	if reverse {
		if op == "<" {
			op = ">="
		} else {
			op = "<="
		}
	}

	s := [2]int{}
	switch op {
	case "<":
		s[0] = 1
		s[1] = val
	case ">":
		s[0] = val + 1
		s[1] = 4001
	case "<=":
		s[0] = 1
		s[1] = val + 1
	case ">=":
		s[0] = val
		s[1] = 4001
	}

	r[idx] = overlapInterval(r[idx], s)

	return r
}

func combs(r intervals) int {
	res := 1
	for _, i := range r {
		a := i[0]
		b := i[1]
		if b < a {
			return 0
		}
		res *= b - a
	}
	return res
}

func createIntervals() intervals {
	res := intervals{}
	for i := 0; i < 4; i++ {
		res[i] = [2]int{1, 4001}
	}
	return res
}

func main() {
	inpB, _ := io.ReadAll(os.Stdin)
	inpS := string(inpB)

	input := strings.Split(inpS, "\n\n")

	wfStr := input[0]
	wfSlice := strings.Split(wfStr, "\n")

	wfMap := map[string][]string{}
	accPos := []mapPos{}
	for _, wf := range wfSlice {
		nameSpl := strings.Split(wf, "{")
		name := nameSpl[0]
		rest := nameSpl[1]
		rest = rest[:len(rest)-1]
		sepWfs := strings.Split(rest, ",")
		for i, swf := range sepWfs {
			if strings.Contains(swf, "A") {
				accPos = append(accPos, mapPos{name, i})
			}
		}
		wfMap[name] = sepWfs
	}

	// // part1
	// partsStr := input[1]
	// partsSlice := strings.Split(partsStr, "\n")

	// parts := []part{}
	// for _, p := range partsSlice {
	// 	q := p[1 : len(p)-1]
	// 	d := strings.Split(q, ",")
	// 	partt := part{}
	// 	for i, dd := range d {
	// 		ddd := dd[2:]
	// 		v, _ := strconv.Atoi(ddd)
	// 		partt[i] = v
	// 	}
	// 	parts = append(parts, partt)
	// }

	// acc := []part{}
	// rej := []part{}
	// for _, p := range parts {
	// 	currWf := wfMap["in"]
	// outerLoop:
	// 	for {
	// 		for i := 0; i < len(currWf); i++ {
	// 			indWf := currWf[i]
	// 			t, ok := wfFunc(p, indWf)
	// 			if ok {
	// 				if t == "A" {
	// 					acc = append(acc, p)
	// 					break outerLoop
	// 				} else if t == "R" {
	// 					rej = append(rej, p)
	// 					break outerLoop
	// 				} else {
	// 					currWf = wfMap[t]
	// 					continue outerLoop
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// res := 0
	// for _, a := range acc {
	// 	for _, b := range a {
	// 		res += b
	// 	}
	// }
	// fmt.Println(res)

	prevMap := map[string]mapPos{}
	for k, _ := range wfMap {
		for kk, vv := range wfMap {
			for i, wf := range vv {
				if getTarget(wf) == k {
					prevMap[k] = mapPos{kk, i}
				}
			}
		}
	}

	accIntervals := []intervals{}
	res := 0
	for _, ap := range accPos {
		r := createIntervals()
		curr := ap
		for {
			r = process(curr, wfMap[curr.name], r)
			if curr.name == "in" {
				break
			}
			curr = prevMap[curr.name]
		}
		accIntervals = append(accIntervals, r)
		res += combs(r)
	}

	fmt.Println(res)
}
