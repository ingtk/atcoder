package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var w = bufio.NewWriter(os.Stdout)

func main() {
	n := pint(input())
	list := make([]bool, (2*n)+1)
	for {
		c := choose(list)
		fmt.Fprintln(w, c)
		w.Flush()
		a := pint(input())
		if a == 0 {
			break
		}
		list[c-1] = true
		list[a-1] = true
	}
}

func choose(list []bool) int {
	for i, v := range list {
		if !v {
			return i + 1
		}
	}
	return 0
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

func pint(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
