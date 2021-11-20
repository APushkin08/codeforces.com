// https://codeforces.com/problemset/problem/617/A
package main

import (
	"fmt"
)

func main() {
	var x, rez int
	fmt.Scan(&x)
	if x%5 != 0 {
		rez = x/5 + 1
	} else {
		rez = x / 5
	}
	fmt.Println(rez)
}
