// https://codeforces.com/problemset/problem/136/A
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)

	var buf int
	for i := 0; i < n; i++ {
		fmt.Scan(&buf)
		arr[buf-1] = i + 1
	}

	for _, v := range arr {
		fmt.Print(v, " ")
	}
}
