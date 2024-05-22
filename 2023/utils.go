package main

import (
	"container/heap"
	"regexp"
	"strconv"
)

func ints(s string) []int {
	p := regexp.MustCompile(`-?\d+`)
	r := []int{}
	for _, e := range p.FindAllString(s, -1) {
		n, _ := strconv.Atoi(e)
		r = append(r, n)
	}
	return r
}

func tokens(s string, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(s, -1)
}

func all[S any](s []S, f func(S) bool) bool {
	for _, e := range s {
		if !f(e) {
			return false
		}
	}
	return true
}

func slicesFilter[S any](ts []S, f func(S) bool) []S {
	us := []S{}
	for _, e := range ts {
		if f(e) {
			us = append(us, e)
		}
	}
	return us
}

func slicesMap[S, T any](ts []S, f func(S) T) []T {
	us := []T{}
	for _, e := range ts {
		us = append(us, f(e))
	}
	return us
}

func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func sum(s ...int) (r int) {
	for _, e := range s {
		r += e
	}
	return
}

func prod(s ...int) int {
	r := 1
	for _, e := range s {
		r *= e
	}
	return r
}

func pow(x, n int) (r int) {
	if n < 0 {
		return 0
	}
	for {
		if n%2 == 1 {
			r *= x
		}
		n /= 2
		if n == 0 {
			break
		}
		x *= x
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mapsKeys[K comparable, V any](m map[K]V) (r []K) {
	for k := range m {
		r = append(r, k)
	}
	return
}

func mapsValues[K comparable, V any](m map[K]V) (r []V) {
	for _, v := range m {
		r = append(r, v)
	}
	return
}

func counter[S comparable](s []S) map[S]int {
	r := map[S]int{}
	for _, e := range s {
		r[e]++
	}
	return r
}

func transpose[S any](m [][]S) [][]S {
	rows := len(m)
	cols := len(m[0])

	r := make([][]S, cols)
	for i := 0; i < cols; i++ {
		r[i] = make([]S, rows)
		for j := 0; j < rows; j++ {
			r[i][j] = m[j][i]
		}
	}
	return r
}

func gridFind[S any](grd [][]S, f func(S) bool) [][2]int {
	var res [][2]int
	for i, r := range grd {
		for j, c := range r {
			if f(c) {
				res = append(res, [2]int{i, j})
			}
		}
	}
	return res
}

type item[S any] struct {
	value    S
	priority int
}

type priorityQ[S any] []*item[S]

func (pq priorityQ[S]) Len() int { return len(pq) }

func (pq priorityQ[S]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityQ[S]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQ[S]) Push(x any) {
	item := x.(*item[S])
	*pq = append(*pq, item)
}

func (pq *priorityQ[S]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func newPQ[S any](items ...*item[S]) priorityQ[S] {
	pq := make(priorityQ[S], 0)
	heap.Init(&pq)
	for _, e := range items {
		heap.Push(&pq, e)
	}
	return pq
}

type pair struct {
	r, c int
}
