// https://codeforces.com/problemset/problem/492/B
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, l int
	fmt.Scan(&n, &l)
	a := make([]int, n)

	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	max := a[0]
	if (l - a[n-1]) > max {
		max = l - a[n-1]
	}
	max *= 2

	for i := 1; i < n; i++ {
		if (a[i] - a[i-1]) > max {
			max = a[i] - a[i-1]
		}
	}

	fmt.Printf("%.10f", float64(max)/float64(2))

}
