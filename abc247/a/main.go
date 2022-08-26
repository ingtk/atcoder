package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	s := input()
	s = "0" + s

	fmt.Println(s[:4])
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
