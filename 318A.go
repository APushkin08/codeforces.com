// https://codeforces.com/problemset/problem/318/A
package main

import (
	"fmt"
)

func main() {
	var n, k uint64
	fmt.Scan(&n, &k)

	if k <= n/2 {
		fmt.Print(k*2 - 1)
	} else if n%2 == 0 {
		fmt.Print((k - (n / 2)) * 2)
	} else if k == ((n + 1) / 2) {
		fmt.Println(n)
	} else {
		fmt.Print((k - ((n + 1) / 2)) * 2)
	}
}
