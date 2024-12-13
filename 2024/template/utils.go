package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

/*
int parsing
*/

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func ints(s string) (r []int) {
	p := regexp.MustCompile(`-?\d+`)
	for _, e := range p.FindAllString(s, -1) {
		r = append(r, atoi(e))
	}
	return
}

func pints(s string) (r []int) {
	p := regexp.MustCompile(`\d+`)
	for _, e := range p.FindAllString(s, -1) {
		r = append(r, atoi(e))
	}
	return
}

/*
matrix parsing
*/

func readMatrix[T any](s string, transform func(byte) T) [][]T {
	rows := strings.Split(s, "\n")
	matrix := make([][]T, len(rows))

	for i, row := range rows {
		matrix[i] = make([]T, len(row))
		for j := range row {
			matrix[i][j] = transform(row[j])
		}
	}

	return matrix
}

/*
linear algebra
*/

type pt [2]int

type vec pt

func (u vec) add(v vec) vec {
	return vec{u[0] + v[0], u[1] + v[1]}
}

func (u vec) mul(c int) vec {
	return vec{c * u[0], c * u[1]}
}

func (u vec) rotate(n int) vec {
	a, b := u[0], u[1]
	for range n % 4 {
		a, b = -b, a
	}
	return vec{a, b}
}

func transpose[T any](m [][]T) [][]T {
	R, C := len(m), len(m[0])
	m2 := make([][]T, C)
	for i := range C {
		m2[i] = make([]T, R)
	}

	for r := range R {
		for c := range C {
			m2[c][r] = m[r][c]
		}
	}

	return m2
}

func rotate[T any](m [][]T, n int) [][]T {
	for range n % 4 {
		m = transpose(m)

		for c := range len(m) {
			slices.Reverse(m[c])
		}
	}
	return m
}

func jordanGauss(A [][]float64, b []float64) ([]float64, bool) {
	mul := func(r []float64, c float64) []float64 {
		res := []float64{}
		for _, e := range r {
			res = append(res, c*e)
		}
		return res
	}

	add := func(r []float64, t []float64) []float64 {
		res := []float64{}
		for i := 0; i < len(r); i++ {
			res = append(res, r[i]+t[i])
		}
		return res
	}

	R := len(A)

	for i := range R {
		A[i] = append(A[i], b[i])
	}

	for i := 0; i < R; i++ {
		A[i] = mul(A[i], 1/A[i][i])
		for j := i + 1; j < R; j++ {
			A[j] = add(A[j], mul(A[i], -A[j][i]))
		}
	}

	for i := R - 1; i >= 0; i-- {
		A[i] = mul(A[i], 1/A[i][i])
		for j := i - 1; j >= 0; j-- {
			A[j] = add(A[j], mul(A[i], -A[j][i]))
		}
	}

	res := []float64{}
	for i := 0; i < R; i++ {
		if math.IsNaN(A[i][R]) || math.IsInf(A[i][R], +1) || math.IsInf(A[i][R], -1) {
			return []float64{}, false
		}
		res = append(res, A[i][R])
	}

	return res, true
}

/*
number theory
*/

func gcdExtended(a, b int) (int, int, int) {
	if a == 0 {
		return b, 0, 1
	}
	gcd, x1, y1 := gcdExtended(b%a, a)
	x := y1 - (b/a)*x1
	y := x1
	return gcd, x, y
}

func crt(nums []int, rems []int) int {
	prod := 1
	for _, n := range nums {
		prod *= n
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		prodI := prod / nums[i]
		_, invI, _ := gcdExtended(prodI, nums[i])
		result += rems[i] * prodI * invI
	}

	return result % prod
}

/*
misc
*/

type set[T comparable] map[T]bool

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func pop[T any](slice *[]T) T {
	n := len(*slice)
	if n == 0 {
		panic("empty slice")
	}
	back := (*slice)[n-1]
	*slice = (*slice)[:n-1]
	return back
}

func popFront[T any](slice *[]T) T {
	if len(*slice) == 0 {
		panic("empty slice")
	}
	front := (*slice)[0]
	*slice = (*slice)[1:]
	return front
}

func reverse(s string) (r string) {
	tmp := []byte(s)
	slices.Reverse(tmp)
	return string(tmp)
}

/*
min heap
*/

type lesser[T any] interface {
	less(other T) bool
}

type item[T any] struct {
	state    T
	distance int
}

func newItem[T any](state T, distance int) *item[T] {
	return &item[T]{state: state, distance: distance}
}

func (p item[T]) less(other item[T]) bool {
	return p.distance < other.distance
}

type minHeap[T lesser[T]] struct {
	arr []*T
}

func newMinHeap[T lesser[T]]() *minHeap[T] {
	return &minHeap[T]{arr: []*T{}}
}

func (h *minHeap[T]) push(value *T) {
	h.arr = append(h.arr, value)
	h.up(h.len() - 1)
}

func (h *minHeap[T]) pop() *T {
	if h.len() == 0 {
		panic("empty heap")
	}
	min_ := h.arr[0]
	n := h.len() - 1
	h.arr[0] = h.arr[n]
	h.arr = h.arr[:n]
	h.down(0, n)
	return min_
}

func (h *minHeap[T]) up(j int) {
	for j > 0 {
		i := (j - 1) / 2
		if i == j || !(*h.arr[j]).less(*h.arr[i]) {
			break
		}
		h.arr[j], h.arr[i] = h.arr[i], h.arr[j]
		j = i
	}
}

func (h *minHeap[T]) down(i0, n int) {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && (*h.arr[j2]).less(*h.arr[j1]) {
			j = j2
		}
		if !(*h.arr[j]).less(*h.arr[j]) {
			break
		}
		h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
		i = j
	}
}

func (h *minHeap[T]) len() int {
	return len(h.arr)
}

/*
performance
*/

func perf[S, T any](f func(S) T, input S) {
	measure := func() int64 {
		t0 := time.Now()
		_ = f(input)
		return time.Since(t0).Microseconds()
	}

	firstRun := measure()
	n := int(1000000.0 / float64(firstRun))
	if n <= 1 || firstRun < 10 {
		fmt.Printf("%.3fms\n", float64(firstRun)/1000.0)
		return
	}

	var runTimes []int64
	for range n {
		runTimes = append(runTimes, measure())
	}

	sum := int64(0)
	for _, t := range runTimes {
		sum += t
	}

	fmt.Printf("%.3f ms\n", (float64(sum)/float64(n))/1000.0)
}
