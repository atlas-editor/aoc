package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y, z int
}

type velocity struct {
	vx, vy, vz int
}

type path struct {
	p point
	v velocity
}

type slopeInterceptForm struct {
	k, q float64
}

func getSlopeInterceptForm(p path) slopeInterceptForm {
	a := p.v.vy
	b := -p.v.vx
	c := -(a*p.p.x + b*p.p.y)

	return slopeInterceptForm{float64(-a) / float64(b), float64(-c) / float64(b)}
}

func intersection(l0, l1 slopeInterceptForm) (float64, float64, bool) {
	a, c := l0.k, l0.q
	b, d := l1.k, l1.q
	if a == b {
		return 0.0, 0.0, false
	}

	return (d - c) / (a - b), (a * (d - c) / (a - b)) + c, true

}

func ptInFuture(pt path, x, y float64) bool {
	diff := x - float64(pt.p.x)
	r := diff / float64(pt.v.vx)

	if r > 0 {
		return true
	}

	return false
}

type row []float64

func (r row) mul(c float64) row {
	res := row{}
	for _, e := range r {
		res = append(res, c*e)
	}
	return res
}

func (r row) add(t row) row {
	res := row{}
	for i := 0; i < len(r); i++ {
		res = append(res, r[i]+t[i])
	}
	return res
}

func (r row) switchLastSign() row {
	r[len(r)-1] = -r[len(r)-1]
	return r
}

func jordanGauss(m []row) []float64 {
	R := len(m)

	for i := 0; i < R; i++ {
		m[i] = m[i].mul(1 / m[i][i])
		for j := i + 1; j < R; j++ {
			m[j] = m[j].add(m[i].mul(-m[j][i]))
		}
	}

	for i := R - 1; i >= 0; i-- {
		m[i] = m[i].mul(1 / m[i][i])
		for j := i - 1; j >= 0; j-- {
			m[j] = m[j].add(m[i].mul(-m[j][i]))
		}
	}

	res := []float64{}
	for i := 0; i < R; i++ {
		res = append(res, m[i][R])
	}
	return res
}

func roundToInts(f []float64) []int {
	res := []int{}
	for _, e := range f {
		res = append(res, int(math.Round(e)))
	}
	return res
}

func sum(nums []int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

func buildEq1(p path) row {
	return row{
		-float64(p.v.vy),
		float64(p.v.vx),
		0.0,
		float64(p.p.y),
		-float64(p.p.x),
		0.0,
		float64(p.p.x)*float64(p.v.vy) - float64(p.p.y)*float64(p.v.vx)}
}

func buildEq2(p path) row {
	return row{
		-float64(p.v.vz),
		0.0,
		float64(p.v.vx),
		float64(p.p.z),
		0.0,
		-float64(p.p.x),
		float64(p.p.x)*float64(p.v.vz) - float64(p.p.z)*float64(p.v.vx)}
}

func main() {
	scn := bufio.NewScanner(os.Stdin)

	paths := []path{}
	anals := []slopeInterceptForm{}
	for scn.Scan() {
		line := scn.Text()
		q := strings.Split(line, "@")

		w := strings.TrimSpace(q[0])
		ww := strings.Split(w, ", ")
		px, _ := strconv.Atoi(strings.TrimSpace(ww[0]))
		py, _ := strconv.Atoi(strings.TrimSpace(ww[1]))
		pz, _ := strconv.Atoi(strings.TrimSpace(ww[2]))

		e := strings.TrimSpace(q[1])
		ee := strings.Split(e, ", ")
		vx, _ := strconv.Atoi(strings.TrimSpace(ee[0]))
		vy, _ := strconv.Atoi(strings.TrimSpace(ee[1]))
		vz, _ := strconv.Atoi(strings.TrimSpace(ee[2]))

		p := path{point{px, py, pz}, velocity{vx, vy, vz}}
		paths = append(paths, p)
		anals = append(anals, getSlopeInterceptForm(p))
	}

	// part1
	B0, B1 := 200000000000000.0, 400000000000000.0
	c := 0
	for i := 0; i < len(anals); i++ {
		for j := i + 1; j < len(anals); j++ {
			x, y, ok := intersection(anals[i], anals[j])
			if ok && x >= B0 &&
				x <= B1 && y >= B0 &&
				y <= B1 &&
				ptInFuture(paths[i], x, y) &&
				ptInFuture(paths[j], x, y) {
				c++
			}
		}
	}
	fmt.Println(c)

	// part2
	eq0XY := buildEq1(paths[0])
	eq0XZ := buildEq2(paths[0])
	equations := []row{}
	for i := 1; i < 4; i++ {
		eqiXY := eq0XY.add(buildEq1(paths[i]).mul(-1)).switchLastSign()
		eqiXZ := eq0XZ.add(buildEq2(paths[i]).mul(-1)).switchLastSign()
		equations = append(equations, eqiXY, eqiXZ)
	}

	fmt.Println(sum(roundToInts(jordanGauss(equations))[:3]))
}
