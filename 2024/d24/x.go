package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func main() {
	path := "input.txt"
	//path = "sample.txt"

	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))

	a := []string{}
	b := []string{}
	for i := range 45 {
		xi := fmt.Sprintf("x%02d", i)
		yi := fmt.Sprintf("y%02d", i)
		a = append(a, xi)
		b = append(b, yi)
	}
	z := add2(a, b)

	for i := range len(a) + 1 {
		fmt.Printf("z%v = %v\n", i, z[i])
	}

	//for k, v := range equations {
	//	fmt.Println(k, v)
	//}

	//fmt.Println(p2(input))
}

var re = regexp.MustCompile(`(.*) (AND|OR|XOR) (.*) -> (.*)`)
var re2 = regexp.MustCompile(`(.*) (AND|OR|XOR) (.*)`)
var equations = map[string]string{}

func p1(input string) int {
	parts := strings.Split(input, "\n\n")

	knownValues := map[string]int{}
	for _, line := range strings.Split(parts[0], "\n") {
		tmp := strings.Split(line, ": ")
		knownValues[tmp[0]] = atoi(tmp[1])
	}

	tmpEquations := strings.Split(parts[1], "\n")
	for _, e := range tmpEquations {
		x := re.FindStringSubmatch(e)
		equations[sortEq(e)] = x[4]
	}

	//for k, v := range equations {
	//	fmt.Println(k, "--", v)
	//}

	//x := re.FindStringSubmatch(tmpEquations[0])
	//
	//for i, e := range x {
	//	fmt.Println(i, e)
	//}

	//for i, e := range tmpEquations {
	//	fmt.Println(i, e)
	//}

	seenVar := set[string]{}
	for k := range knownValues {
		seenVar[k] = true
	}

	seenEq := set[string]{}
	q := []string{}
	for _, e := range tmpEquations {
		tmp := re.FindStringSubmatch(e)
		input0, input1, output := tmp[1], tmp[3], tmp[4]
		if seenVar[input0] && seenVar[input1] {
			q = append(q, e)
			seenEq[e] = true
			seenVar[output] = true
			break
		}
	}

	p := []string{}
	for len(q) > 0 {
		curr := popFront(&q)
		p = append(p, curr)

		for _, e := range tmpEquations {
			tmp := re.FindStringSubmatch(e)
			input0, input1, output := tmp[1], tmp[3], tmp[4]
			if !seenEq[e] && seenVar[input0] && seenVar[input1] {
				q = append(q, e)
				seenEq[e] = true
				seenVar[output] = true
			}
		}
	}

	//fmt.Println(p)

	//slices.SortFunc(tmpEquations, func(a, b string) int {
	//	tmp := re.FindStringSubmatch(a)
	//	aInput := []string{tmp[1], tmp[3]}
	//	aOutput := tmp[4]
	//
	//	tmp = re.FindStringSubmatch(b)
	//	bInput := []string{tmp[1], tmp[3]}
	//	bOutput := tmp[4]
	//
	//	if slices.Contains(bInput, aOutput) {
	//		return -1
	//	}
	//	if slices.Contains(aInput, bOutput) {
	//		return 1
	//	}
	//	return 0
	//})

	for _, e := range p {
		//fmt.Println(e)
		tmp := re.FindStringSubmatch(e)
		input0, op, input1, output := tmp[1], tmp[2], tmp[3], tmp[4]
		switch op {
		case "AND":
			knownValues[output] = knownValues[input0] & knownValues[input1]
		case "OR":
			knownValues[output] = knownValues[input0] | knownValues[input1]
		case "XOR":
			knownValues[output] = knownValues[input0] ^ knownValues[input1]
		default:
			panic("unknown op")
		}
	}

	//fmt.Println(knownValues)

	return -1
}

func p2(input string) int {
	return -1
}

func and(x, y []string) []string {
	z := []string{}
	for i := range min(len(x), len(y)) {
		zz := fmt.Sprintf("%v AND %v", x[i], y[i])
		zi := "(" + zz + ")"
		if len(x[i]) == 3 && len(y[i]) == 3 {
			if v, ok := equations[sortedEq(zz)]; ok {
				zi = v
			} else {
				zi = "(" + sortedEq(zz) + ")"
			}
		}
		if x[i] == "0" || y[i] == "0" {
			zi = "0"
		}
		z = append(z, zi)
	}
	for i := min(len(x), len(y)); i < max(len(x), len(y)); i++ {
		z = append(z, "0")
	}

	return z
}

func xor(x, y []string) []string {
	z := []string{}
	for i := range min(len(x), len(y)) {
		zz := fmt.Sprintf("%v XOR %v", x[i], y[i])
		zi := "(" + zz + ")"
		if len(x[i]) == 3 && len(y[i]) == 3 {
			if v, ok := equations[sortedEq(zz)]; ok {
				zi = v
			} else {
				zi = "(" + sortedEq(zz) + ")"
			}
		}
		if x[i] == "0" && y[i] == "0" {
			zi = "1"
		} else if x[i] == "0" {
			zi = y[i]
		} else if y[i] == "0" {
			zi = x[i]
		}
		z = append(z, zi)
	}
	w := x
	if len(x) < len(y) {
		w = y
	}
	for i := min(len(x), len(y)); i < max(len(x), len(y)); i++ {
		z = append(z, w[i])
	}

	return z
}

func shift(x []string) []string {
	return append([]string{"0"}, x...)
}

func add(a, b []string) []string {
	n := len(a) + 1
	for !isZero(b, n) {
		carry := and(a, b)
		a = xor(a, b)
		b = shift(carry)
	}

	return a
}

func isZero(a []string, n int) bool {
	for i, e := range a {
		if i > n {
			return true
		}
		if e != "0" {
			return false
		}
	}
	return true
}

func sortEq(e string) string {
	x := re.FindStringSubmatch(e)
	y := []string{x[1], x[3]}
	slices.Sort(y)
	return fmt.Sprintf("%v %v %v", y[0], x[2], y[1])
}

func sortedEq(e string) string {
	x := re2.FindStringSubmatch(e)
	y := []string{x[1], x[3]}
	slices.Sort(y)
	return fmt.Sprintf("%v %v %v", y[0], x[2], y[1])
}

//def add_bitwise(x, y):
//	n = len(x)
//	sum_result = [0] * n  # Initialize the sum array
//	carry = 0  # Initialize carry
//
//	for i in range(n):
//	# Calculate sum for the current bit
//	sum_result[i] = x[i] ^ y[i] ^ carry
//	# Calculate new carry
//	carry = (x[i] & y[i]) | (carry & (x[i] ^ y[i]))
//
//	# If there's a carry left after the last bit, handle it
//	if carry:
//	print("Overflow: carry remains after addition")
//
//	return sum_result

func add2(a, b []string) []string {
	n := len(a)
	result := make([]string, n+1)
	carry := "0"

	for i := range n {
		sumBit := equations[sortedEq(fmt.Sprintf("%v XOR %v", a[i], b[i]))]

		if carry != "0" {
			sumBit = equations[sortedEq(fmt.Sprintf("%v XOR %v", sumBit, carry))]
		}

		result[i] = sumBit

		tmp0 := equations[sortedEq(fmt.Sprintf("%v AND %v", a[i], b[i]))]
		tmp1 := equations[sortedEq(fmt.Sprintf("%v XOR %v", a[i], b[i]))]

		if carry != "0" {
			tmp2 := equations[sortedEq(fmt.Sprintf("%v AND %v", carry, tmp1))]
			carry = equations[sortedEq(fmt.Sprintf("%v OR %v", tmp0, tmp2))]
		} else {
			carry = tmp0
		}
	}

	result[n] = carry

	return result
}

// z08 <-> mvb
// jss <-> rds
// wss <-> z18
// bmn <-> z23
