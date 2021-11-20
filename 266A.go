// https://codeforces.com/problemset/problem/266/A
package main

import (
	"fmt"
)

func main() {
	var n, sum int
	var st string
	fmt.Scan(&n)
	fmt.Scan(&st)
	for i := 0; i < n-1; i++ {
		k := i + 1
		if st[i] == st[k] {
			sum++
		}
	}

	fmt.Print(sum)
}
