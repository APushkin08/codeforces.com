// https://codeforces.com/problemset/problem/25/A
package main

import (
	"fmt"
)

func main() {
	var n, parityCount int
	fmt.Scan(&n)
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < 3; i++ {
		if (arr[i] % 2) == 0 {
			parityCount++
		}
	}

	if parityCount >= 2 {
		for i, v := range arr {
			if v%2 != 0 {
				fmt.Println(i + 1)
				break
			}
		}
	} else {
		for i, v := range arr {
			if v%2 == 0 {
				fmt.Println(i + 1)
				break
			}
		}
	}

}
