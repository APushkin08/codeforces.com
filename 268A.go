// https://codeforces.com/problemset/problem/268/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, cnt int
	fmt.Scan(&n)
	h := make([]int, n)
	a := make([]int, n)
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &h[i], &a[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && h[i] == a[j] {
				cnt++
			}
		}
	}

	fmt.Println(cnt)
}
