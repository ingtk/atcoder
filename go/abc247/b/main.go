package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type name struct {
	s string
	t string
}

func main() {
	n := ninput()
	names := make([]name, 0, n)
	for i := 0; i < n; i++ {
		l := inputs(" ")
		n := name{l[0], l[1]}
		names = append(names, n)
	}

	for i := 0; i < n; i++ {
		var sf, tf bool
		si, ti := names[i].s, names[i].t
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			sj, tj := names[j].s, names[j].t
			if si == sj || si == tj {
				sf = true
			}
			if ti == tj || ti == sj {
				tf = true
			}
		}
		if sf && tf {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
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

func pint(s string) int {
	n, _ := strconv.ParseInt(s, 10, 64)
	return int(n)
}
