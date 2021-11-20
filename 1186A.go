// https://codeforces.com/problemset/problem/1186/A
package main

import (
	"fmt"
)

func main() {
	var n, m, k int

	fmt.Scan(&n, &m, &k)

	if n <= m && n <= k {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
