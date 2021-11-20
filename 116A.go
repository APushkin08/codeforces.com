// https://codeforces.com/problemset/problem/116/A
package main

import (
	"fmt"
)

func main() {
	var n, a, b, max, buf int
	fmt.Scan(&n)
	fmt.Scan(&a, &b)
	max, buf = b, b
	for i := 2; i <= n; i++ {
		fmt.Scan(&a, &b)
		buf = buf - a + b
		if max < buf {
			max = buf
		}
	}

	fmt.Print(max)
}
