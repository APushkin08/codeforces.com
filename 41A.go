// https://codeforces.com/problemset/problem/41/A
package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	isConversely := "YES"
	if len(s) != len(t) {
		fmt.Println("NO")

	} else {
		for i, j := 0, len(s)-1; i < len(s); i, j = i+1, j-1 {
			if s[i] != t[j] {
				isConversely = "NO"
			}
		}
		fmt.Println(isConversely)
	}

}
