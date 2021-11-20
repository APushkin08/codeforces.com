// https://codeforces.com/problemset/problem/141/A
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var st1, st2, st3 string
	fmt.Scan(&st1, &st2, &st3)

	st1 = st1 + st2
	mp1 := make(map[rune]int)
	mp2 := make(map[rune]int)

	for _, v := range st1 {
		mp1[v]++
	}

	for _, v := range st3 {
		mp2[v]++
	}

	if reflect.DeepEqual(mp1, mp2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
