package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Shopping struct {
	couponNum   int
	couponValue int
	items       []int
}

func (s *Shopping) Len() int {
	return len(s.items)
}

func (s *Shopping) Less(i, j int) bool {
	return s.items[i]%s.couponValue > s.items[j]%s.couponValue
}

func (s *Shopping) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}

func (s *Shopping) solve() int {
	for i := range s.items {
		j := 0
		for ; j < s.couponNum; j++ {
			if s.items[i]-(s.couponValue*(j+1)) < 0 {
				break
			}
		}
		s.items[i] -= (s.couponValue * j)
		s.couponNum -= j
		if s.couponNum == 0 {
			break
		}
	}

	if s.couponNum > 0 {
		for i := 0; i < len(s.items); i++ {
			s.items[i] -= s.couponValue
			s.couponNum--
			if s.couponNum == 0 {
				break
			}
		}
	}

	total := 0
	for _, v := range s.items {
		if v <= 0 {
			continue
		}
		total += v
	}

	return total
}

func main() {
	l := ninputs(" ")
	a := ninputs(" ")
	k, x := l[1], l[2]
	fmt.Println(solve(k, x, a))
}

//-----------------------------------
func solve(k, x int, a []int) int {
	s := &Shopping{k, x, a}
	sort.Sort(s)
	return s.solve()
}

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func inputs(sep string) []string {
	sc.Scan()
	return strings.Split(sc.Text(), sep)
}

func ninputs(sep string) []int {
	ws := inputs(sep)
	res := make([]int, 0, len(ws))
	for _, v := range ws {
		res = append(res, pint(v))
	}
	return res
}

func pint(s string) int {
	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n)
}
