// https://codeforces.com/problemset/problem/61/A
package main

import (
	"fmt"
)

func main() {
	var st1, st2, rez string
	fmt.Scan(&st1, &st2)

	for i := 0; i < len(st1); i++ {
		if st1[i] != st2[i] {
			rez += "1"
		} else {
			rez += "0"
		}
	}

	fmt.Println(rez)

}
