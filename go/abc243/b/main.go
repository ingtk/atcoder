package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	_ = input("")
	line2 := input(" ")
	line3 := input(" ")
	ans1 := 0
	ans2 := 0
	for i := 0; i < len(line2); i++ {
		for j := 0; j < len(line3); j++ {
			a := pint(line2[i])
			b := pint(line3[j])
			if a == b {
				if i == j {
					ans1++
				} else {
					ans2++
				}
			}
		}
	}

	fmt.Println(ans1)
	fmt.Println(ans2)
}

var scanner = bufio.NewScanner(os.Stdin)

func input(sep string) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), sep)
}

func pint(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}
