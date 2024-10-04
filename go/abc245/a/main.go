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
	l := ninputs(" ")
	a, b, c, d := l[0], l[1], l[2], l[3]

	ans := "Takahashi"
	if (a > c) || (a == c && b > d) {
		ans = "Aoki"
	}
	fmt.Println(ans)
}

//-----------------------------------

var sc *bufio.Scanner

func init() {
	sc = bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func inputs(sep string) []string {
	sc.Scan()
	return strings.Split(sc.Text(), sep)
}

func ninputs(sep string) []int64 {
	ws := inputs(sep)
	res := make([]int64, 0, len(ws))
	for _, v := range ws {
		res = append(res, pint(v))
	}
	return res
}

func pint(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
