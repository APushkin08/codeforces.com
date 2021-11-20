// https://codeforces.com/problemset/problem/133/A
package main

import "fmt"

func main() {
	var input string
	rez := "NO"
	fmt.Scan(&input)
	for _, ch := range input {
		if ch == 'H' || ch == 'Q' || ch == '9' {
			rez = "YES"
			break
		}
	}
	fmt.Println(rez)
}
