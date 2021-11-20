// https://codeforces.com/problemset/problem/231/A
package main

import (
	"fmt"
)

func main() {
	var n, a, b, c, rez, sum int
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b, &c)
		rez += a + b + c
		if rez >= 2 {
			sum++
		}
		rez = 0
	}
	fmt.Print(sum)
}
