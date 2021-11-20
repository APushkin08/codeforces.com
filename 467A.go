// https://codeforces.com/problemset/problem/467/A
package main

import (
	"fmt"
)

func main() {
	var n, p, q int
	var count int

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&p, &q)
		if q-p >= 2 {
			count++
		}
	}
	fmt.Println(count)
}
