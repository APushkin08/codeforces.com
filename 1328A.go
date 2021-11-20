// https://codeforces.com/problemset/problem/1328/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t int
	var a, b int64
	fmt.Scan(&t)

	in := bufio.NewReader(os.Stdin)
	for ; t != 0; t-- {
		fmt.Fscan(in, &a, &b)
		if a%b == 0 {
			fmt.Println(0)
		} else {
			fmt.Println(b - a%b)
		}
	}

}
