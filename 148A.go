// https://codeforces.com/problemset/problem/148/A
package main

import (
	"fmt"
)

func main() {
	var mas [4]int
	var d, cntD int

	for i := 0; i < 4; i++ {
		fmt.Scan(&mas[i])
	}
	fmt.Scan(&d)

	for i := 1; i <= d; i++ {
		for j := 0; j < 4; j++ {
			if i%mas[j] == 0 {
				cntD++
				break
			}
		}
	}

	//	fmt.Print(mas)
	fmt.Print(cntD)
}
