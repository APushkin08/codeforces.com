// https://codeforces.com/problemset/problem/705/A
package main

import (
	"fmt"
)

func main() {
	var n, k uint
	var st string
	k = 1
	fmt.Scan(&n)
	for k <= n {
		if k%2 == 0 {
			st += "I love that "
		} else {
			st += "I hate that "
		}

		k++
	}
	st = st[:len(st)-5] + "it"
	fmt.Println(st)
}
