// https://codeforces.com/problemset/problem/996/A
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(minKol(n))
}

func minKol(sum int) int {
	bills := [...]int{1, 5, 10, 20, 100}
	mode := sum
	rez := 0
	i := 4

	for mode != 0 { //двигаемся по номиналам купюр от самой большой(100) до мелкой(1)
		rez = rez + mode/bills[i] //к результату добавляем количесво купюр текущего номинала
		mode = sum % bills[i]
		//mode = mode % bills[i]
		i-- //переходим к номиналу помельче
	}
	return rez
}
