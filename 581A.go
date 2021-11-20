// https://codeforces.com/problemset/problem/581/A
package main

import (
	"fmt"
)

func main() {
	var r, b, same, diff int
	fmt.Scan(&r, &b)
	diff = min(r, b)
	same = (max(r, b) - diff) / 2

	fmt.Println(diff, same)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
