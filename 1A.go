// https://codeforces.com/problemset/problem/1/A
package main

import "fmt"

func main() {
	var n, m, a int64
	fmt.Scan(&n, &m, &a)
	fmt.Println(((a + m - 1) / a) * ((a + n - 1) / a))

}
