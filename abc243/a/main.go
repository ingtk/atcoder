package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	answers := []string{"F", "M", "T"}
	line := input(" ")
	v := pint(line[0])
	abc := []int64{pint(line[1]), pint(line[2]), pint(line[3])}

	i := 0
	for ; ; i++ {
		v = v - abc[i%3]
		if v < 0 {
			break
		}
	}
	fmt.Println(answers[i%3])
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
