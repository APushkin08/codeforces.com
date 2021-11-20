// https://codeforces.com/problemset/problem/155/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, val, rez, min, max int
	fmt.Scan(&n)

	if n == 1 {
		fmt.Println(0)
	} else {
		fmt.Scan(&min)
		max = min

		in := bufio.NewReader(os.Stdin)
		for i := 2; i <= n; i++ {
			fmt.Fscan(in, &val)
			if val > max {
				max = val
				rez++
				continue
			}
			if val < min {
				min = val
				rez++
			}
		}
		fmt.Println(rez)
	}
}
