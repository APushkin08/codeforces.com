// https://codeforces.com/problemset/problem/486/A
package main

import (
	"fmt"
)

func main() {
	var n, rez int64
	fmt.Scan(&n)

	if n%2 == 0 {
		rez = n / 2
	} else {
		rez = -(n + 1) / 2
	}

	fmt.Println(rez)

}
