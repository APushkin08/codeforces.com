// https://codeforces.com/problemset/problem/489/C
package main

import "fmt"

func main() {
	var m, s int
	fmt.Scan(&m, &s)

	if (s < 1) && (m > 1) || (s > m*9) {
		fmt.Println("-1 -1")
		return
	}

	for i, k := m-1, s; i >= 0; i-- {
		j := max(0, k-9*i)
		if (j == 0) && (i == m-1) && (k != 0) {
			j = 1
		}
		fmt.Print(j)
		k -= j
	}

	fmt.Print(" ")

	for i, k := m-1, s; i >= 0; i-- {
		j := min(9, k)
		fmt.Print(j)
		k -= j
	}

}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
