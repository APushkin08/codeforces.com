// https://codeforces.com/problemset/problem/344/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, mag, rez int
	fmt.Scan(&n)

	buf := -1
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &mag)

		if mag != buf {
			rez++
			buf = mag
		}

	}

	fmt.Println(rez)

}
