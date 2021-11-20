// https://codeforces.com/problemset/problem/228/A
package main

import (
	"fmt"
)

func main() {
	m := make(map[int64]int)

	var in int64
	for i := 0; i < 4; i++ {
		fmt.Scan(&in)
		m[in]++
	}

	fmt.Println(4 - len(m))

}
