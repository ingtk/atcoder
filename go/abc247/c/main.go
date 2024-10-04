package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
	n := ninput()
	fmt.Println(solve(n))
}

func solve(n int) string {
	s := make([][]int, n)
	s[0] = []int{1}

	for i := 1; i < n; i++ {
		s[i] = append(s[i], s[i-1]...)
		s[i] = append(s[i], i+1)
		s[i] = append(s[i], s[i-1]...)
	}

	result := ""

	for _, v := range s[n-1] {
		result += fmt.Sprint(v, " ")
	}

	result = result[:len(result)-1]

	return result
}

//-----------------------------------

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func input() string {
	sc.Scan()
	return sc.Text()
}

func ninput() int {
	return pint(input())
}

func pint(s string) int {
	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n)
}
