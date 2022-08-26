package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type data struct {
	x, c int
}

func main() {
	n := ninput()
	stack := make([]data, 200000)
	ans := []int{}
	for i := 0; i < n; i++ {
		q := ninputs(" ")
		stack, ans = solve(n, q, stack, ans)
	}
	for _, a := range ans {
		fmt.Println(a)
	}
}

func solve(n int, q []int, stack []data, ans []int) ([]data, []int) {
	switch q[0] {
	case 1:
		x, c := q[1], q[2]
		stack = append(stack, data{x, c})
	case 2:
		c := q[1]
		sum := 0
		for {
			d := stack[0]
			if d.c < c {
				sum += d.x * d.c
				c -= d.c
				if len(stack) > 1 {
					stack = stack[1:]
				} else {
					break
				}
			} else {
				sum += d.x * c
				d.c -= c
				c = 0
				stack[0] = d
			}
			if c <= 0 {
				break
			}
		}
		ans = append(ans, sum)
	}

	return stack, ans
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

func inputs(sep string) []string {
	sc.Scan()
	return strings.Split(sc.Text(), sep)
}

func ninput() int {
	return pint(input())
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
