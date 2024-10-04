package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	l1 := ninputs(" ")
	n, k := l1[0], l1[1]
	a := ninputs(" ")
	b := ninputs(" ")

	fmt.Println(solve(n, k, a, b))
}

func solve(n, k int, a, b []int) string {
	dp := map[int]struct{}{}
	dp[a[0]] = struct{}{}
	if a[0] != b[0] {
		dp[b[0]] = struct{}{}
	}

	for i := 1; i < n; i++ {
		xdp := map[int]struct{}{}
		ai := a[i]
		bi := b[i]
		for v := range dp {
			if abs(v-ai) <= k {
				xdp[ai] = struct{}{}
			}
			if ai == bi {
				continue
			}
			if abs(v-bi) <= k {
				xdp[bi] = struct{}{}
			}
		}
		dp = map[int]struct{}{}
		for i := range xdp {
			dp[i] = struct{}{}
		}
	}

	if len(dp) == 0 {
		return "No"
	}
	return "Yes"
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return n * -1
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
