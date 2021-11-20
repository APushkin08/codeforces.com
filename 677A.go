// https://codeforces.com/problemset/problem/677/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, h int
	var rez int
	fmt.Scan(&n, &h)

	var inBuf int
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &inBuf)

		if inBuf <= h {
			rez++
		} else {
			rez += 2
		}
	}

	fmt.Println(rez)
}
