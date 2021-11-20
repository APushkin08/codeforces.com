// https://codeforces.com/problemset/problem/469/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, x, y int
	var rez string
	fmt.Scan(&n)
	sl := make([]int, n)

	fmt.Scan(&x)

	var num int
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < x; i++ {
		fmt.Fscan(in, &num)
		sl[num-1]++
	}

	fmt.Scan(&y)
	for i := 0; i < y; i++ {
		fmt.Fscan(in, &num)
		sl[num-1]++
	}

	rez = "I become the guy."
	for _, v := range sl {
		if v == 0 {
			rez = "Oh, my keyboard!"
		}
	}

	fmt.Println(rez)

}
