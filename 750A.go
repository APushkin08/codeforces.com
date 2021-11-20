// https://codeforces.com/problemset/problem/750/A
package main

import (
	"fmt"
)

func main() {
	var n, k, t, res int
	fmt.Scan(&n, &k)

	for i := 1; i <= n; i++ {
		t += i * 5
		if t <= 240-k {
			res = i
		} else {
			break
		}
	}
	fmt.Println(res)
}
