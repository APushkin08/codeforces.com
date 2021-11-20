// http://codeforces.com/problemset/problem/455/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var mx int64
	fmt.Scan(&n)
	var f [100005]int64
	var freq [100005]int64

	in := bufio.NewReader(os.Stdin)
	var ai int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &ai)
		freq[ai]++
		mx = max(int64(mx), int64(ai))
	}

	f[1] = freq[1]
	for i := 2; int64(i) <= mx; i++ {
		f[i] = max(f[i-2]+(freq[i]*int64(i)), f[i-1])
	}

	fmt.Println(f[mx])
}

func max(n1, n2 int64) int64 {
	if n1 > n2 {
		return n1
	}
	return n2
}
