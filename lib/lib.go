package lib

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func ninput() int64 {
	return pint(input())
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
