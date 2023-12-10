package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func gcd(a, b *big.Int) *big.Int {
	if b.Cmp(big.NewInt(0)) == 0 {
		return a
	}
	r := new(big.Int)
	r.Mod(a, b)
	return gcd(b, r)
}

func lcm(a, b *big.Int) *big.Int {
	p := new(big.Int)
	p.Mul(a, b)
	gcdVal := gcd(a, b)
	return p.Div(p, gcdVal)
}

func lcmOfSlice(nums []int) *big.Int {
	res := big.NewInt(int64(nums[0]))

	for i := 1; i < len(nums); i++ {
		res = lcm(res, big.NewInt(int64(nums[i])))
	}

	return res
}

func mapToDest(s []string, m map[string][]string, d rune) []string {
	res := []string{}
	idx := 1
	if d == 'L' {
		idx = 0
	}

	for _, el := range s {
		res = append(res, m[el][idx])
	}

	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	lrs := scanner.Text()
	scanner.Scan()

	instructions := map[string][]string{}
	pts := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		src, b, _ := strings.Cut(line, " = ")
		destL, destR, _ := strings.Cut(b[1:len(b)-1], ", ")
		instructions[src] = []string{destL, destR}
		pts = append(pts, src)
	}

	dests := []string{}
	for _, pt := range pts {
		if pt[len(pt)-1] == 'A' {
			dests = append(dests, pt)
		}
	}

	steps := 0
	lengths := []int{-1, -1, -1, -1, -1, -1}
	for {
		for _, r := range lrs {
			dests = mapToDest(dests, instructions, r)
			steps++
			for i := 0; i < len(dests); i++ {
				if lengths[i] == -1 && dests[i][len(dests[i])-1] == 'Z' {
					lengths[i] = steps
				}
			}
		}
		done := true
		for _, i := range lengths {
			if i == -1 {
				done = false
			}
		}
		if done {
			fmt.Println(lcmOfSlice(lengths))
			break
		}
	}
}
