// https://codeforces.com/problemset/problem/520/A
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var st string
	fmt.Scan(&n)
	fmt.Scan(&st)
	st = strings.ToLower(st)

	mp := make(map[rune]int)
	for _, v := range st {
		mp[v]++
	}

	if len(mp) == 26 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
