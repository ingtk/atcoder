package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/trees/avltree"
)

func main() {
	n := pint(input())
	tr := avltree.NewWithIntComparator()
	for i := 0; i < n; i++ {
		q := ninputs()
		solve(tr, q)
	}
}

func solve(tr *avltree.Tree, q []int) {
	switch q[0] {
	case 1:
		x := q[1]
		v, found := tr.Get(x)
		if found {
			tr.Put(x, v.(int)+1)
		} else {
			tr.Put(x, 1)
		}
	case 2:
		x := q[1]
		k := q[2]
		node, _ := tr.Floor(x)
		if node == nil{
			fmt.Println(-1)
			return
		}
		if k < node.Value.(int) {
			fmt.Println(node.Key)
			return
		}
		cnt := k
		for {
			cnt -= node.Value.(int)
			if cnt <= 0 {
				break
			}
			node = node.Prev()
			if node == nil {
				fmt.Println(-1)
				return
			}
		}
		fmt.Println(node.Key)
	case 3:
		x := q[1]
		k := q[2]
		node, _ := tr.Ceiling(x)
		if node == nil {
			fmt.Println(-1)
			return
		}
		if k < node.Value.(int) {
			fmt.Println(node.Key)
			return
		}
		cnt := k
		for {
			cnt -= node.Value.(int)
			if cnt <= 0 {
				break
			}
			node = node.Next()
			if node == nil {
				fmt.Println(-1)
				return
			}
		}
		fmt.Println(node.Key)
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

func inputs(seps ...string) []string {
	sep := " "
	if len(seps) > 0 {
		sep = seps[0]
	}
	sc.Scan()
	return strings.Split(sc.Text(), sep)
}

func ninput() int {
	return pint(input())
}

func ninputs(seps ...string) []int {
	ws := inputs(seps...)
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
