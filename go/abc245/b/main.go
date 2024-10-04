package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Numbers []int64

func main() {
	_ = input()
	l := ninputs(" ")
	m := map[int]struct{}{}

	for _, v := range l {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}

	for i := 0; i <= 2000; i++ {
		if _, ok := m[i]; !ok {
			fmt.Println(i)
			break
		}
	}
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
