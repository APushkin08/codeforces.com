// https://codeforces.com/problemset/problem/959/A
package main

import (
	"fmt"
)

func main() {
	var n uint32

	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Mahmoud")
	} else {
		fmt.Println("Ehab")
	}

}
