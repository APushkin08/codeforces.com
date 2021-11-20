// https://codeforces.com/problemset/problem/734/A
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var st string
	fmt.Scan(&n, &st)

	cntA := strings.Count(st, "A")

	if cntA > len(st)-cntA {
		fmt.Println("Anton")
	} else if cntA < len(st)-cntA {
		fmt.Println("Danik")
	} else {
		fmt.Println("Friendship")
	}

}
