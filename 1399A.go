// https://codeforces.com/problemset/problem/1399/A
package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)

	for t != 0 {
		var n int
		fmt.Scan(&n)
		arr := make([]int, n)
		for i := range arr {
			fmt.Scan(&arr[i])
		}

		sort.Ints(arr)

		var fl bool = true
		for j := 1; j < n; j++ {
			fl = fl && (arr[j]-arr[j-1] <= 1)
		}

		if fl == true {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

		t--
	}
}
