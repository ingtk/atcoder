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
	_ = input()
	s := input()
	x, y := solve(s)
	fmt.Printf("%d %d\n", x, y)
}

func solve(s string) (int, int) {
	dirs := []int{0, 0, 0, 0}
	curDir := 1
	for _, v := range s {
		switch v {
		case 'R':
			curDir = (curDir + 1) % 4
		case 'S':
			dirs[curDir]++
		}
	}
	y := dirs[0] - dirs[2]
	x := dirs[1] - dirs[3]

	return x, y
}

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
