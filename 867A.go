// https://codeforces.com/problemset/problem/867/A
package main

import "fmt"

func main() {
	var n int
	var st string

	fmt.Scan(&n)
	fmt.Scan(&st)
	if string([]rune(st)[0]) == "S" && string([]rune(st)[len(st)-1]) == "F" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	//fmt.Println(string([]rune(st)[len(st)-1]))

}
