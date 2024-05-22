//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	res := 1
	for scanner.Scan() {
		line := scanner.Text()
		nums := ints(line)

		bpNo, oreRobotCostOre, calyRobotCostOre, obisidanRobotCostOre, obisidanRobotCostClay, geodeRobotCostOre, geodeRobotCostObsidian := nums[0], nums[1], nums[2], nums[3], nums[4], nums[5], nums[6]

		var dp func(minutes, oreRobots, clayRobots, obisidianRobots, geodeRobots, ore, clay, obsidian, geode int) int
		cache := map[[9]int]int{}
		dp = func(minutes, oreRobots, clayRobots, obisidianRobots, geodeRobots, ore, clay, obsidian, geode int) int {
			if val, found := cache[[9]int{minutes, oreRobots, clayRobots, obisidianRobots, geodeRobots, ore, clay, obsidian, geode}]; found {
				return val
			}

			if minutes == 0 {
				return geode
			}

			r0 := dp(minutes-1, oreRobots, clayRobots, obisidianRobots, geodeRobots, ore+oreRobots, clay+clayRobots, obsidian+obisidianRobots, geode+geodeRobots)
			r1, r2, r3, r4 := 0, 0, 0, 0
			if ore >= oreRobotCostOre {
				r1 = dp(minutes-1, oreRobots+1, clayRobots, obisidianRobots, geodeRobots, ore-oreRobotCostOre+oreRobots, clay+clayRobots, obsidian+obisidianRobots, geode+geodeRobots)
			}
			if ore >= calyRobotCostOre {
				r2 = dp(minutes-1, oreRobots, clayRobots+1, obisidianRobots, geodeRobots, ore-calyRobotCostOre+oreRobots, clay+clayRobots, obsidian+obisidianRobots, geode+geodeRobots)
			}
			if ore >= obisidanRobotCostOre && clay >= obisidanRobotCostClay {
				r3 = dp(minutes-1, oreRobots, clayRobots, obisidianRobots+1, geodeRobots, ore-obisidanRobotCostOre+oreRobots, clay-obisidanRobotCostClay+clayRobots, obsidian+obisidianRobots, geode+geodeRobots)
			}
			if ore >= geodeRobotCostOre && obsidian >= geodeRobotCostObsidian {
				r4 = dp(minutes-1, oreRobots, clayRobots, obisidianRobots, geodeRobots+1, ore-geodeRobotCostOre+oreRobots, clay+clayRobots, obsidian-geodeRobotCostObsidian+obisidianRobots, geode+geodeRobots)
			}

			cache[[9]int{minutes, oreRobots, clayRobots, obisidianRobots, geodeRobots, ore, clay, obsidian, geode}] = max(r0, r1, r2, r3, r4)
			return max(r0, r1, r2, r3, r4)
		}

		best := dp(32, 1, 0, 0, 0, 0, 0, 0, 0)

		fmt.Println(bpNo, best)
		res *= best
	}
	fmt.Println("\n", res)
}
