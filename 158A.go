// https://codeforces.com/problemset/problem/158/A
package main

import (
	"fmt"
)

func main() {
	var n, k, summ int

	fmt.Scan(&n, &k)

	mass := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&mass[i])
	}

	for i := 0; i < n; i++ {
		if mass[i] > 0 && mass[i] >= mass[k-1] {
			summ++
		}
	}

	fmt.Println(summ)

}
