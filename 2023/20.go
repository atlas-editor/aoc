package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"slices"
	"strings"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, nums ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(nums); i++ {
		result = LCM(result, nums[i])
	}

	return result
}

func LCMSlice(nums []int) int {
	return LCM(nums[0], nums[1], nums[2:]...)
}

type module struct {
	name         string
	mType        string
	targets      []string
	on           bool
	dependencies map[string]string
}

type msg struct {
	m      string
	signal string
}

func allHigh(moduleName string, mm map[string]module) bool {
	for _, v := range mm[moduleName].dependencies {
		if v != "high" {
			return false
		}
	}
	return true
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	moduleMap := map[string]module{}
	for scn.Scan() {
		line := scn.Text()
		inOut := strings.Split(line, " -> ")
		in, out := inOut[0], inOut[1]
		targets := strings.Split(out, ", ")

		currModule := module{}
		currModule.targets = targets
		if in == "broadcaster" {
			currModule.name = in
			moduleMap[currModule.name] = currModule
			continue
		}

		currModule.name = in[1:]
		currModule.mType = in[:1]

		if currModule.mType == "%" {
			currModule.on = false
		}

		currModule.dependencies = map[string]string{}

		moduleMap[currModule.name] = currModule
	}

	for _, v := range moduleMap {
		if v.mType == "&" {
			for _, vv := range moduleMap {
				if slices.Contains(vv.targets, v.name) {
					v.dependencies[vv.name] = "low"
				}
			}
		}
	}

	resHigh := 0
	resLow := 0
	cycles := map[string]int{}
	i := 1
	// for part1 the for cycle has to loop 1000 times
	for {
		highPulses := 0
		lowPulses := 1
		q := list.New()
		q.PushBack(msg{"broadcaster", "low"})
		for q.Len() > 0 {
			val := q.Front()
			q.Remove(val)
			curr := val.Value.(msg)
			currModule := moduleMap[curr.m]
			currSignal := curr.signal
			if currModule.mType == "&" {
				if allHigh(curr.m, moduleMap) {
					currSignal = "low"
				} else {
					currSignal = "high"
				}
				// from input file lg is the only module sending pulses to rx
				if curr.m == "lg" {
					if len(currModule.dependencies) == len(cycles) {
						nums := []int{}
						for _, v := range cycles {
							nums = append(nums, v)
						}
						fmt.Println(LCMSlice(nums))
						return
					}
					for k, v := range currModule.dependencies {
						if _, found := cycles[k]; !found && v == "high" {
							cycles[k] = i
						}
					}
				}
			}

			if currSignal == "high" {
				highPulses += len(currModule.targets)
			} else {
				lowPulses += len(currModule.targets)
			}

			for _, t := range currModule.targets {
				currTarget := moduleMap[t]
				if currTarget.mType == "%" && currSignal == "low" {
					if currTarget.on {
						q.PushBack(msg{currTarget.name, "low"})
					} else {
						q.PushBack(msg{currTarget.name, "high"})
					}
					currTarget.on = !currTarget.on
					moduleMap[t] = currTarget
				}

				if currTarget.mType == "&" {
					currTarget.dependencies[currModule.name] = currSignal
					moduleMap[t] = currTarget
					q.PushBack(msg{currTarget.name, "TBD"})
				}
			}
		}
		resHigh += highPulses
		resLow += lowPulses
		i++
	}
	// part 1
	// fmt.Println(resHigh * resLow)

}
