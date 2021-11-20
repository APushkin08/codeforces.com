// https://codeforces.com/problemset/problem/935/A
package main

import (
	"fmt"
)

func main() {
	var n, rez int

	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			rez++
		}
	}

	fmt.Println(rez)
}
