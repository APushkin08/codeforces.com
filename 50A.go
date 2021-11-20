// https://codeforces.com/problemset/problem/50/A
package main

import (
	"fmt"
)

func main() {
	var m, n, rez int
	fmt.Scanln(&m, &n)
	rez = (m * n) / 2
	fmt.Print(rez)
}
