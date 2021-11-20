// https://codeforces.com/problemset/problem/977/A
package main

import (
	"fmt"
	"os"
)

func main() {
	var num1 uint32
	var num2 uint8

	fmt.Fscan(os.Stdin, &num1)
	fmt.Scan(&num2)

	var i uint8 = 1

	for ; i <= num2; i++ {

		if (num1 % 10) == 0 {
			num1 /= 10
		} else {
			num1--
		}
	}
	fmt.Println(num1)

}
