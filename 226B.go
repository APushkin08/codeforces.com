// https://codeforces.com/problemset/problem/266/B
package main

import (
	"fmt"
	"strings"
)

func main() {
	var st string
	var n, t int
	fmt.Scan(&n, &t)
	fmt.Scan(&st)

	for i := 0; i < t; i++ {
		st = strings.ReplaceAll(st, "BG", "GB")
	}

	fmt.Println(st)
}
