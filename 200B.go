// https://codeforces.com/problemset/problem/200/B
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, num, sum int
	var rez float64

	fmt.Scan(&n)

	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &num)
		sum += num
	}

	rez = float64(sum) / float64(n)

	fmt.Printf("%.12f", rez)

}
