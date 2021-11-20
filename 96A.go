// https://codeforces.com/problemset/problem/96/A
package main

import (
	"fmt"
	s "strings"
)

func main() {
	var st string
	fmt.Scan(&st)
	if s.Contains(st, "1111111") || s.Contains(st, "0000000") {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
