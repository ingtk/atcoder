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
	line := inputs(" ")
	x, n := pint(line[0]), pint(line[1])
	s := input()
	fmt.Println(solve(x, n, s))
}

func solve(n, x int64, s string) int64 {
	ss := shorten(s)
	if len(ss) == 0 {
		return x
	}

	for i := 0; i < len(ss); i++ {
		switch ss[i] {
		case 'L':
			x = (x * 2)
		case 'R':
			x = (x * 2) + 1
		case 'U':
			x = x / 2
		}
	}
	return x
}

func shorten(s string) []rune {
	stack := make([]rune, 0, len(s))
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		switch v {
		case 'U':
			lv := stack[len(stack)-1]
			if lv == 'L' || lv == 'R' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		case 'L', 'R':
			stack = append(stack, v)
		}
	}
	return stack
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
