// https://codeforces.com/problemset/problem/282/A
package main

import (
	"fmt"
)

func main() {
	var n, rez int
	var st string

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&st)
		if string([]rune(st)[1]) == "+" {
			rez++
		} else {
			rez--
		}
	}

	fmt.Print(rez)
}
