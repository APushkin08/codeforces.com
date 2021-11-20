// https://codeforces.com/problemset/problem/110/A
package main

import (
	"fmt"
	"strings"
)

func main() {
	var st string
	var cnt int
	fmt.Scan(&st)

	cnt += strings.Count(st, "4")
	cnt += strings.Count(st, "7")

	if cnt == 7 || cnt == 4 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
