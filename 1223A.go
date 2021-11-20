// https://codeforces.com/problemset/problem/1223/A
package main

import (
	"fmt"
)

func main() {
	var q uint32
	var n uint32
	var sl []uint32

	fmt.Scan(&q)

	for i := 1; uint32(i) <= q; i++ {
		fmt.Scan(&n)
		if n == 2 {
			sl = append(sl, 2)
			//fmt.Println(2)
		} else {
			sl = append(sl, n&1)
			//fmt.Println(n & 1) //1 - если четно, иначе 0
		}
	}
	for _, val := range sl {
		fmt.Println(val)
	}

}
