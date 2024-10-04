package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
}

func solve(n, m int) {
}

//-----------------------------------

type sortInt []int

func (n sortInt) Len() int           { return len(n) }
func (n sortInt) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n sortInt) Less(i, j int) bool { return n[i] < n[j] }

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func input() string {
	sc.Scan()
	return sc.Text()
}

func inputs(seps ...string) []string {
	sep := " "
	if len(seps) > 0 {
		sep = seps[0]
	}
	sc.Scan()
	return strings.Split(sc.Text(), sep)
}

func ninput() int {
	return pint(input())
}

func ninputs(seps ...string) []int {
	ws := inputs(seps...)
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
