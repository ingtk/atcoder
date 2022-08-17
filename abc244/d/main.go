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
	s := inputs(" ")
	t := inputs(" ")

	diffCnt := 0
	for i := 0; i < 3; i++ {
		if s[i] != t[i] {
			diffCnt++
		}
	}
	switch diffCnt {
	case 0, 3:
		fmt.Println("Yes")
	default:
		fmt.Println("No")
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

func pint(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
