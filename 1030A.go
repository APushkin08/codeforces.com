// https://codeforces.com/problemset/problem/1030/A
package main

import (
	"fmt"
)

func main() {
	var n int
	var flag uint8
	var easyFlag bool = true

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		fmt.Scan(&flag)
		if flag == 1 {
			easyFlag = false
			break
		}

	}

	if easyFlag {
		fmt.Println("Easy")
	} else {
		fmt.Println("Hard")
	}

}
