package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// // part1
// var cardOrder = "AKQJT98765432"
// part2
var cardOrder = "AKQT98765432J"

type Hand struct {
	cards string
	bid   int
}

func handRank(h Hand) (int, []int) {
	cards := map[rune]int{}
	for _, r := range h.cards {
		cards[r]++
	}
	reps := []int{}
	for ch, v := range cards {
		if ch == 'J' {
			continue
		}
		reps = append(reps, v)
	}
	slices.Sort(reps)
	slices.Reverse(reps)

	if len(reps) != 0 {
		reps[0] += cards['J']
	} else {
		reps = append(reps, cards['J'])
	}

	r0 := -1
	if reps[0] == 5 {
		r0 = 0
	} else if reps[0] == 4 {
		r0 = 1
	} else if reps[0] == 3 {
		if reps[1] == 2 {
			r0 = 2
		} else {
			r0 = 3
		}
	} else if reps[0] == 2 {
		if reps[1] == 2 {
			r0 = 4
		} else {
			r0 = 5
		}
	} else {
		r0 = 6
	}

	r1 := []int{}
	for _, ch := range h.cards {
		r1 = append(r1, strings.Index(cardOrder, string(ch)))
	}

	return r0, r1
}

func handCmp(h0, h1 Hand) int {
	h0r0, h0r1 := handRank(h0)
	h1r0, h1r1 := handRank(h1)
	if h0r0 != h1r0 {
		return h0r0 - h1r0
	}

	return slices.Compare(h0r1, h1r1)
}

func counter(h string) map[string]int {
	unique := map[string]int{}
	for _, r := range h {
		unique[string(r)]++
	}
	return unique
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	hs := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		lineFields := strings.Fields(line)
		hand := lineFields[0]
		bid, _ := strconv.Atoi(lineFields[1])
		h := Hand{hand, bid}
		hs = append(hs, h)
	}

	slices.SortFunc(hs, handCmp)

	sum := 0
	l := len(hs)
	for i, h := range hs {
		sum += (l - i) * h.bid
	}
	fmt.Println(sum)
}
