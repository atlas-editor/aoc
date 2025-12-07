//go:build ignore

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]
	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	deps := map[string][]string{}
	val := map[string]int{}
	op := map[string]string{}
	for _, line := range strings.Split(input, "\n") {
		a := strings.Split(line, ": ")
		lhs := a[0]
		b := strings.Fields(a[1])
		if len(b) == 1 {
			val[lhs] = atoi(b[0])
		} else {
			deps[lhs] = []string{b[0], b[2]}
			op[lhs] = b[1]
		}
	}

	q := []string{"root"}
	order := []string{}
	for len(q) > 0 {
		// fmt.Println(q)
		curr := pop(&q)
		order = append(order, curr)
		for _, n := range deps[curr] {
			q = append(q, n)
		}
	}

	expr := map[string]*expression{}

	for k, v := range val {
		expr[k] = &expression{op: "const", val: v}
	}
	expr["humn"] = &expression{op: "var"}

	for i := len(order) - 1; i >= 0; i-- {
		m := order[i]
		if _, ok := expr[m]; !ok {
			expr[m] = &expression{op: op[m], lhs: expr[deps[m][0]], rhs: expr[deps[m][1]]}
		}
	}

	rhsVal := expr[deps["root"][1]].simplify().val
	lhs := expr[deps["root"][0]].simplify()

	curr := lhs
	for {
		if curr.op == "var" {
			break
		}
		if curr.rhs.op == "const" {
			switch curr.op {
			case "+":
				rhsVal -= curr.rhs.val
			case "-":
				rhsVal += curr.rhs.val
			case "*":
				rhsVal /= curr.rhs.val
			case "/":
				rhsVal *= curr.rhs.val
			}
			curr = curr.lhs
		} else if curr.lhs.op == "const" {
			switch curr.op {
			case "+":
				rhsVal -= curr.lhs.val
			case "-":
				rhsVal = -rhsVal + curr.lhs.val
			case "*":
				rhsVal /= curr.lhs.val
			}
			curr = curr.rhs
		} else {
			break
		}
	}

	fmt.Println(rhsVal)
}

type expression struct {
	op  string
	lhs *expression
	rhs *expression
	val int
}

func (e *expression) simplify() *expression {
	if e.op == "const" || e.op == "var" {
		return e
	}

	lhs := e.lhs.simplify()
	rhs := e.rhs.simplify()

	if lhs.op == "const" && rhs.op == "const" {
		switch e.op {
		case "+":
			return &expression{op: "const", val: lhs.val + rhs.val}
		case "-":
			return &expression{op: "const", val: lhs.val - rhs.val}
		case "*":
			return &expression{op: "const", val: lhs.val * rhs.val}
		case "/":
			return &expression{op: "const", val: lhs.val / rhs.val}
		}
	}

	return &expression{op: e.op, lhs: lhs, rhs: rhs}
}

func (e *expression) String() string {
	if e.op == "const" {
		return strconv.Itoa(e.val)
	}

	switch e.op {
	case "const":
		return strconv.Itoa(e.val)
	case "var":
		return "X"
	default:
		return fmt.Sprintf("(%v %v %v)", e.lhs.String(), e.op, e.rhs.String())
	}

}

func pop[T any](slice *[]T) T {
	n := len(*slice)
	if n == 0 {
		panic("empty slice")
	}
	first := (*slice)[0]
	*slice = (*slice)[1:]
	return first
}

func atoi(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}
