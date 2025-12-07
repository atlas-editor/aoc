//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	data int
	next *node
	prev *node
}

func (n *node) move(i int) {
	dest := n
	for range i {
		dest = dest.next
	}
	for range -i + 1 {
		dest = dest.prev
	}

	n.prev.next = n.next
	n.next.prev = n.prev

	n.next = dest.next
	dest.next.prev = n
	dest.next = n
	n.prev = dest
}

func (n *node) print() {
	fmt.Printf("-> %v ", n.data)

	curr := n.next
	for curr.data != n.data {
		fmt.Printf("-> %v ", curr.data)
		curr = curr.next
	}
	fmt.Println("->")
}

func (n *node) grove() int {
	zero := n
	for zero.data != 0 {
		zero = zero.next
	}

	num := zero
	s := 0
	for i := range 3000 {
		num = num.next
		if i == 999 || i == 1999 || i == 2999 {
			// fmt.Println(num.data)
			s += num.data
		}
	}
	return s
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func p(nodes []*node, loops int, key int) int {
	for _, n := range nodes {
		n.data *= key
	}

	for range loops {
		for _, n := range nodes {
			n.move(n.data % (len(nodes) - 1))
		}
	}
	return nodes[0].grove()
}

func main() {
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	nodes := []*node{}
	for _, line := range strings.Split(input, "\n") {
		nodes = append(nodes, &node{data: atoi(line)})
	}

	for i := 1; i < len(nodes)-1; i++ {
		a, b, c := nodes[i-1], nodes[i], nodes[i+1]
		b.prev = a
		b.next = c
	}
	nodes[0].next = nodes[1]
	nodes[0].prev = nodes[len(nodes)-1]
	nodes[len(nodes)-1].next = nodes[0]
	nodes[len(nodes)-1].prev = nodes[len(nodes)-2]

	// fmt.Println(p(nodes, 1, 1))
	fmt.Println(p(nodes, 10, 811589153))
}
