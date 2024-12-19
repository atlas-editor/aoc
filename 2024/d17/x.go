package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]

	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(input string) string {
	parts := strings.Split(input, "\n\n")
	registers := make([][]int, 3)
	for i, line := range strings.Split(parts[0], "\n") {
		registers[i] = ints(line)
	}

	A, B, C := registers[0][0], registers[1][0], registers[2][0]
	program := ints(parts[1])

	output := []string{}
	for _, n := range simulate(program, 0, A, B, C) {
		output = append(output, strconv.Itoa(n))
	}

	return strings.Join(output, ",")
}

func p2(input string) int {
	parts := strings.Split(input, "\n\n")
	registers := make([][]int, 3)
	for i, line := range strings.Split(parts[0], "\n") {
		registers[i] = ints(line)
	}

	program := ints(parts[1])
	return solve(program)
}

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	base := x
	for range n - 1 {
		x *= base
	}
	return x
}

func combo(operand, A, B, C int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	default:
		panic(fmt.Sprintf("invalid combo operand: %v", operand))
	}
}

func simulate(program []int, i, A, B, C int) []int {
	idx := 0

	output := []int{}
	for {
		if i >= len(program) {
			break
		}
		opcode := program[i]
		operand := program[i+1]

		switch opcode {
		case 0:
			a := A
			b := pow(2, combo(operand, A, B, C))
			A = a / b
			//A = A >> 3
		case 1:
			a := B
			b := operand
			B = a ^ b
			//B = B ^ 0b011
		case 2:
			a := combo(operand, A, B, C)
			b := 8
			B = a % b
			//B = A & 0b111
		case 3:
			if A != 0 {
				i = operand
				//i = 0
				continue
			}
		case 4:
			a := B
			b := C
			B = a ^ b
			//B = B ^ C
		case 5:
			a := combo(operand, A, B, C)
			b := 8
			out := a % b
			//out := B & 0b111
			output = append(output, out)
			idx++
		case 6:
			a := A
			b := pow(2, combo(operand, A, B, C))
			B = a / b
			//panic("unused opcode")
		case 7:
			a := A
			b := pow(2, combo(operand, A, B, C))
			C = a / b
			//C = A >> B
		default:
			panic(fmt.Sprintf("invalid opcode: %v", opcode))
		}

		i += 2
	}

	return output
}

func solve(target []int) int {
	var f func([]int, int, []int) (int, bool)
	f = func(rem []int, i int, acc []int) (int, bool) {
		if len(rem) == 0 || i == len(target) {
			bin := ""
			for _, digit := range append(rem, acc...) {
				if digit < 0 {
					bin += "0"
					continue
				}
				bin += strconv.Itoa(digit)
			}
			A, _ := strconv.ParseInt(bin, 2, 0)
			if slices.Equal(simulate(target, 0, int(A), 0, 0), target) {
				return int(A), true
			}
			return -1, false
		}
		for _, p := range fill(rem, target[i]) {
			if v, ok := f(p[:len(p)-3], i+1, append(p[len(p)-3:], acc...)); ok {
				return v, true
			}
		}

		return -1, false
	}

	nums := make([]int, 64)
	for i := range 64 {
		nums[i] = -1
	}

	val, _ := f(nums, 0, []int{})
	return val
}

func fill(num []int, out int) [][]int {
	nn := []int{}
	for j := range 10 {
		if j >= len(num) {
			break
		}

		k := len(num) - 1 - j
		nn = append(nn, num[k])
	}

	advanced := [][]int{}
	outBin := toBin(out)
	for B := range 8 {
		n := slices.Clone(nn)
		BBin := toBin(B)
		if (n[0] == -1 || n[0] == BBin[2]) && (n[1] == -1 || n[1] == BBin[1]) && (n[2] == -1 || n[2] == BBin[0]) {
			n[0] = BBin[2]
			n[1] = BBin[1]
			n[2] = BBin[0]
		} else {
			continue
		}

		switch B {
		case 0:
			if len(n) > 5 && (n[5] == -1 || n[5] == outBin[0]) && (n[4] == -1 || n[4] == outBin[1]) && (n[3] == -1 || n[3] == outBin[2]) {
				n[5] = outBin[0]
				n[4] = outBin[1]
				n[3] = outBin[2]

				advNum := slices.Clone(num)
				for i := range 6 {
					k := len(advNum) - 1 - i
					advNum[k] = n[i]
				}
				advanced = append(advanced, advNum)
			}
		case 1:
			if len(n) > 4 && (n[4] == -1 || n[4] == outBin[0]) && (n[3] == -1 || n[3] == outBin[1]) && 1 == outBin[2] {
				n[4] = outBin[0]
				n[3] = outBin[1]

				advNum := slices.Clone(num)
				for i := range 5 {
					k := len(advNum) - 1 - i
					advNum[k] = n[i]
				}
				advanced = append(advanced, advNum)
			}
		case 2:
			if len(n) > 3 && (n[3] == -1 || n[3] == outBin[0]) && 1 == outBin[1] && 1 == outBin[2] {
				n[3] = outBin[0]

				advNum := slices.Clone(num)
				for i := range 4 {
					k := len(advNum) - 1 - i
					advNum[k] = n[i]
				}
				advanced = append(advanced, advNum)
			}
		case 3:
			if 0 == outBin[0] && 0 == outBin[1] && 0 == outBin[2] {
				advNum := slices.Clone(num)
				for i := range 3 {
					k := len(advNum) - 1 - i
					advNum[k] = n[i]
				}
				advanced = append(advanced, advNum)
			}
		case 4:
			if len(n) > 9 && (n[9] == -1 || n[9] == 1-outBin[0]) && (n[8] == -1 || n[8] == outBin[1]) && (n[7] == -1 || n[7] == outBin[2]) {
				n[9] = 1 - outBin[0]
				n[8] = outBin[1]
				n[7] = outBin[2]

				advNum := slices.Clone(num)
				for i := range 10 {
					k := len(advNum) - 1 - i
					advNum[k] = n[i]
				}
				advanced = append(advanced, advNum)

			}
		case 5:
			if len(n) > 8 && (n[8] == -1 || n[8] == 1-outBin[0]) && (n[7] == -1 || n[7] == outBin[1]) && (n[6] == -1 || n[6] == 1-outBin[2]) {
				n[8] = 1 - outBin[0]
				n[7] = outBin[1]
				n[6] = 1 - outBin[2]

				advNum := slices.Clone(num)
				for i := range 9 {
					k := len(advNum) - 1 - i
					advNum[k] = n[i]
				}
				advanced = append(advanced, advNum)
			}
		case 6:
			if len(n) > 7 && (n[7] == -1 || n[7] == 1-outBin[0]) && (n[6] == -1 || n[6] == 1-outBin[1]) && (n[5] == -1 || n[5] == outBin[2]) {
				n[7] = 1 - outBin[0]
				n[6] = 1 - outBin[1]
				n[5] = outBin[2]

				advNum := slices.Clone(num)
				for i := range 8 {
					k := len(advNum) - 1 - i
					advNum[k] = n[i]
				}
				advanced = append(advanced, advNum)
			}
		case 7:
			if len(n) > 6 && (n[6] == -1 || n[6] == 1-outBin[0]) && (n[5] == -1 || n[5] == 1-outBin[1]) && (n[4] == -1 || n[4] == 1-outBin[2]) {
				n[6] = 1 - outBin[0]
				n[5] = 1 - outBin[1]
				n[4] = 1 - outBin[2]

				advNum := slices.Clone(num)
				for i := range 7 {
					k := len(advNum) - 1 - i
					advNum[k] = n[i]
				}
				advanced = append(advanced, advNum)
			}
		}
	}

	return advanced
}

func toBin(n int) [3]int {
	b := [3]int{}
	k := strconv.FormatInt(int64(n), 2)
	i := 2
	j := len(k) - 1
	for j >= 0 {
		b[i] = atoi(string(k[j]))
		j--
		i--
	}
	return b
}

/*
utils
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
