// https://codeforces.com/problemset/problem/112/A
package main

import (
	"fmt"
	s "strings"
)

func main() {
	var a, b string
	var rez int

	fmt.Scan(&a, &b)
	a = s.ToUpper(a)
	b = s.ToUpper(b)

	for i := 0; i < len(a); i++ {
		if string([]rune(a)[i]) == string([]rune(b)[i]) {
			continue
		} else if string([]rune(a)[i]) > string([]rune(b)[i]) {
			rez = 1
			break
		} else {
			rez = -1
			break
		}
	}

	fmt.Printf("\n%d", rez)
}
