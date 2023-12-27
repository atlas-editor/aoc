package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Item struct {
	value    vertex
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.index)
}

type vertex string

type cut struct {
	X, Y []vertex
}

type edge struct {
	u, v vertex
	w    int
}

type graph map[vertex]map[vertex]int

func (g graph) addEdge(u, v vertex, w int) {
	if g[u] == nil {
		g[u] = make(map[vertex]int)
	}
	if g[v] == nil {
		g[v] = make(map[vertex]int)
	}
	g[u][v] = w
	g[v][u] = w
}

func (g graph) addOrIncreaseEdge(u, v vertex, w int) {
	if g[u] == nil {
		g[u] = make(map[vertex]int)
	}
	if g[v] == nil {
		g[v] = make(map[vertex]int)
	}
	g[u][v] += w
	g[v][u] += w
}

func (g graph) getVertex() vertex {
	for v := range g {
		return v
	}
	panic("empty graph")
}

func (g graph) removeVertex(u vertex) {
	delete(g, u)
	keys := make([]vertex, 0, len(g))
	for k := range g {
		keys = append(keys, k)
	}

	for _, k := range keys {
		if _, found := g[k][u]; found {
			delete(g[k], u)
		}
	}
}

func (g graph) shrink(u, v vertex) {
	newVertex := vertex(u + "^" + v)
	g[newVertex] = map[vertex]int{}
	for nbr, val := range g[u] {
		g.addOrIncreaseEdge(newVertex, nbr, val)
	}
	for nbr, val := range g[v] {
		if nbr == newVertex {
			continue
		}
		g.addOrIncreaseEdge(newVertex, nbr, val)
	}
	g.removeVertex(u)
	g.removeVertex(v)
}

func (g graph) minimumCutPhase(a vertex) ([2]vertex, int) {
	ordering := []vertex{a}

	pq := make(PriorityQueue, len(g)-1)
	pqPosition := map[vertex]*Item{}
	i := 0
	for v := range g {
		if v == a {
			continue
		}
		pq[i] = &Item{v, g[a][v], i}
		pqPosition[v] = pq[i]
		i++
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		v := item.value
		ordering = append(ordering, v)
		for u, val := range g[v] {
			if _, found := pqPosition[u]; found {
				pq.update(pqPosition[u], pqPosition[u].priority+val)
			}
		}
		delete(pqPosition, v)
	}
	s, t := ordering[len(ordering)-2], ordering[len(ordering)-1]
	cotp := 0
	for _, v := range g[t] {
		cotp += v
	}

	return [2]vertex{s, t}, cotp
}

// stoer-wagner algorithm
func (g graph) minimumCut() cut {
	originalVertices := make([]vertex, 0, len(g))
	for v := range g {
		originalVertices = append(originalVertices, v)
	}

	minCutVal := -1
	minCutX := vertex("")
	for len(g) > 1 {
		a := g.getVertex()
		st, cotp := g.minimumCutPhase(a)
		s, t := st[0], st[1]
		if minCutVal < 0 || minCutVal > cotp {
			minCutVal = cotp
			minCutX = t
		}
		g.shrink(s, t)
	}
	X := []vertex{}
	for _, v := range strings.Split(string(minCutX), "^") {
		X = append(X, vertex(v))
	}
	Y := []vertex{}
	for _, v := range originalVertices {
		if !slices.Contains(X, v) {
			Y = append(Y, v)
		}
	}
	return cut{X, Y}
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	G := graph{}
	for scn.Scan() {
		line := scn.Text()
		data := strings.Split(line, ": ")
		u := data[0]
		nbrs := strings.Fields(data[1])

		for _, v := range nbrs {
			G.addEdge(vertex(u), vertex(v), 1)
		}
	}
	minCut := G.minimumCut()
	fmt.Println(len(minCut.X) * len(minCut.Y))
}
