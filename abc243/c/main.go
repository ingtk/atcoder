package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	x, y int64
	d    int
}

const (
	R = 0
	L = 1
)

type Persons []Person

func (ps Persons) Len() int           { return len(ps) }
func (ps Persons) Less(i, j int) bool { return ps[i].y < ps[j].y }
func (ps Persons) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

func main() {
	n := pint(input())
	persons := make(Persons, 0, n)
	for i := int64(0); i < n; i++ {
		line := inputs(" ")
		p := Person{pint(line[0]), pint(line[1]), 0}
		persons = append(persons, p)
	}
	s := input()

	fmt.Println(solve(n, persons, s))
}

func solve(n int64, persons Persons, s string) string {
	personMap := map[int64]Persons{}
	for i, p := range persons {
		v, ok := personMap[p.y]

		d := R
		if string(s[i]) == "L" {
			d = L
		}

		p.d = d
		if ok {
			personMap[p.y] = append(v, p)
		} else {
			personMap[p.y] = Persons{p}
		}
	}

	ans := "No"

	for _, persons := range personMap {
		if len(persons) <= 1 {
			continue
		}

		var rx int64 = math.MaxInt64
		var lx int64 = math.MinInt64
		for _, p := range persons {
			switch p.d {
			case R:
				if rx > p.x {
					rx = p.x
				}
			case L:
				if lx < p.x {
					lx = p.x
				}
			}
		}
		if rx < lx {
			ans = "Yes"
			break
		}
	}
	return ans
}

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func input() string {
	sc.Scan()
	return sc.Text()
}

func inputs(sep string) []string {
	sc.Scan()
	return strings.Split(sc.Text(), sep)
}

func pint(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
