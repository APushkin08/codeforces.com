// https://codeforces.com/problemset/problem/281/A
package main

import (
	"fmt"
	s "strings"
)

func main() {
	var st string

	fmt.Scanln(&st)

	//s.ToUpper(stingst[1]),
	fmt.Print(s.ToUpper(string(st[0])), st[1:len(st)])

}
