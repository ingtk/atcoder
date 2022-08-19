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
	l2 := ninputs(" ")
	l3 := ninputs(" ")

	x, y := 0, 0
	if l2[0] == l3[0] {
		x = l1[0]
	}
	if l1[0] == l3[0] {
		x = l2[0]
	}
	if l1[0] == l2[0] {
		x = l3[0]
	}
	if l2[1] == l3[1] {
		y = l1[1]
	}
	if l1[1] == l3[1] {
		y = l2[1]
	}
	if l1[1] == l2[1] {
		y = l3[1]
	}

	fmt.Println(x, y)
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
