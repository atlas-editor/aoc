//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

type pair struct {
	x, y int
}

// [x,y)
type intervals []pair

func (i intervals) cmpIntervals(p0, p1 pair) int {
	return p0.y - p1.x
}

func (i intervals) add(xy pair) intervals {
	ibsCmp := func(p pair, n int) int {
		if n < p.x {
			return 1
		} else if n > p.y {
			return -1
		}
		return 0
	}
	xIdx, xFound := slices.BinarySearchFunc(i, xy.x, ibsCmp)
	yIdx, yFound := slices.BinarySearchFunc(i, xy.y, ibsCmp)

	newIntervalX := xy.x
	newIntervalY := xy.y
	if xFound {
		newIntervalX = i[xIdx].x
	}
	if yFound {
		newIntervalY = i[yIdx].y
		yIdx++
	}

	return slices.Replace(i, xIdx, yIdx, pair{newIntervalX, newIntervalY})
}

func (i intervals) cover(xy pair) (bool, int) {
	for _, i0 := range i {
		if i0.x <= xy.x && i0.y >= xy.y {
			return true, -1
		} else if i0.x > xy.x && i0.x < xy.y {
			return false, i0.x - 1
		} else if i0.y < xy.y && i0.y > xy.x {
			return false, i0.y
		}
	}

	return false, -1
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	blocked := map[int]intervals{}
	for scn.Scan() {
		line := scn.Text()

		coordinates := ints(line)
		sensor := pair{coordinates[0], coordinates[1]}
		beacon := pair{coordinates[2], coordinates[3]}

		radius := abs(sensor.x-beacon.x) + abs(sensor.y-beacon.y)
		for i := sensor.y - radius; i <= sensor.y+radius; i++ {
			diff := radius - abs(sensor.y-i)
			if sensor.x-diff > sensor.x+diff+1 {
				fmt.Println(sensor, beacon, diff)
				return
			}
			currInterval := pair{sensor.x - diff, sensor.x + diff + 1}
			blocked[i] = blocked[i].add(currInterval)
		}
	}

	limit := 4000001
	for i := 0; i < limit; i++ {
		if c, idx := blocked[i].cover(pair{0, limit}); !c {
			fmt.Println(idx*4000000 + i)
			return
		}
	}
}
