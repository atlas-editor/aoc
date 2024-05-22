//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type vertex string

type edge struct {
	u, v vertex
	w    int
}

type graph map[vertex]map[vertex]bool

func (g graph) addEdge(u, v vertex) {
	if g[u] == nil {
		g[u] = make(map[vertex]bool)
	}
	if g[v] == nil {
		g[v] = make(map[vertex]bool)
	}
	g[u][v] = true
	g[v][u] = true
}

func floydWarshall(g graph) map[vertex]map[vertex]int {
	d := map[vertex]map[vertex]int{}
	for u := range g {
		d[u] = map[vertex]int{}
	}
	for u := range g {
		d[u] = map[vertex]int{}
		for v := range g {
			if g[u][v] {
				d[u][v] = 1
				d[v][u] = 1
			} else if u == v {
				d[u][u] = 0
			} else {
				d[u][v] = 1000000
				d[v][u] = 1000000
			}
		}
	}

	for u := range g {
		for v := range g {
			for w := range g {
				if d[v][w] > d[u][v]+d[u][w] {
					d[v][w] = d[u][v] + d[u][w]
					d[w][v] = d[u][v] + d[u][w]
				}
			}
		}
	}

	return d
}

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

type item struct {
	t0   int
	u    vertex
	open string
}

type item2 struct {
	t0, t1 int
	u, v   vertex
	open   string
}

type item3 struct {
	v vertex
	t int
}

func mapsKeys(m map[vertex]bool) string {
	ss := []string{}
	for v, t := range m {
		if t {
			ss = append(ss, string(v))
		}
	}
	slices.Sort(ss)
	return strings.Join(ss, "")
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	valveRate := map[string]int{}
	g := graph{}
	for scn.Scan() {
		line := scn.Text()
		valves := tokens(line, `[A-Z]{2}`)
		currValve := valves[0]
		valveRate[currValve] = ints(line)[0]
		for _, v := range valves[1:] {
			g.addEdge(vertex(currValve), vertex(v))
		}
	}

	fw := floydWarshall(g)

	var left func(graph, map[vertex]bool) []vertex
	left = func(g graph, open map[vertex]bool) []vertex {
		r := []vertex{}
		for v := range g {
			if !open[v] && valveRate[string(v)] != 0 {
				r = append(r, v)
			}
		}
		return r
	}

	var dp func(t int, u vertex, open map[vertex]bool) map[item]int
	// cache := map[item]int{}
	dp = func(t int, u vertex, open map[vertex]bool) map[item]int {
		dpTable := map[item]int{}

		for t := 0; t <= 30; t++ {
			for u := range g {
				for vi := range g {
					if valveRate[string(vi)] == 0 || u == vi || t-fw[u][vi]-1 <= 0 {
						continue
					}

					open := make(map[vertex]bool)
					open[vi] = true
					curr := dpTable[item{t - fw[u][vi] - 1, vi, mapsKeys(open)}] + valveRate[string(vi)]*(t-fw[u][vi]-1)
					open[vi] = false

					if curr > dpTable[item{t, u, mapsKeys(open)}] {
						dpTable[item{t, u, mapsKeys(open)}] = curr
					}
				}
			}
		}

		return dpTable
		// for vi := range g {
		// 	if valveRate[string(vi)] == 0 || open[vi] || u == vi || t-fw[u][vi]-1 <= 0 {
		// 		continue
		// 	}

		// 	open[vi] = true
		// 	curr := dp(t-fw[u][vi]-1, vi, open, allowed) + valveRate[string(vi)]*(t-fw[u][vi]-1)
		// 	open[vi] = false
		// 	res = max(curr, res)
		// }
		// cache[item{t, u, mapsKeys(open)}] = res
		// return res
	}

	var dp2 func(t0, t1 int, u, v vertex, open map[vertex]bool) int
	cache2 := map[item2]int{}
	dp2 = func(t0, t1 int, u, v vertex, open map[vertex]bool) int {
		fmt.Println(len(cache2))
		if t0 <= 0 && t1 <= 0 {
			return 0
		}
		if val, found := cache2[item2{t0, t1, u, v, mapsKeys(open)}]; found {
			return val
		}
		var res int

		for _, vi := range left(g, open) {
			for _, vj := range left(g, open) {
				if vi == vj || u == vi || v == vj {
					continue
				}

				if t0-fw[u][vi]-1 <= 0 && t1-fw[v][vj]-1 <= 0 {
					continue
				}

				open[vi] = true
				open[vj] = true

				val0 := valveRate[string(vi)] * (t0 - fw[u][vi] - 1)
				val1 := valveRate[string(vj)] * (t1 - fw[v][vj] - 1)
				newT0 := max(t0-fw[u][vi]-1, 0)
				newT1 := max(t1-fw[v][vj]-1, 0)

				if t0-fw[u][vi]-1 <= 0 {
					val0 = 0
				}
				if t1-fw[v][vj]-1 <= 0 {
					val1 = 0
				}

				curr := dp2(newT0, newT1, vi, vj, open) + val0 + val1

				open[vi] = false
				open[vj] = false
				res = max(curr, res)
			}
		}

		cache2[item2{t0, t1, u, v, mapsKeys(open)}] = res
		return res
	}
	// fmt.Println(dp2(26, 26, vertex("AA"), vertex("AA"), map[vertex]bool{}))
	fmt.Println(dp(30, vertex("AA"), map[vertex]bool{}))

	// currVertex := item3{vertex("AA"), 26}

	// for currVertex.v != vertex("") {
	// 	fmt.Println(currVertex)
	// 	currVertex = path[currVertex]
	// }
}
