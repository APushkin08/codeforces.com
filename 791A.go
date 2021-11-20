// https://codeforces.com/problemset/problem/791/A
package main

import (
	"fmt"
)

func main() {
	var a, b uint
	var year uint = 0

	fmt.Scan(&a, &b)
	for a <= b {
		a *= 3
		b *= 2
		year++
	}
	fmt.Println(year)
}
