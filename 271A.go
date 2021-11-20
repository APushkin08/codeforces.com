// https://codeforces.com/problemset/problem/271/A
package main

import "fmt"

func main() {
	var y int
	fmt.Scan(&y)
	y++
	for !isDif(y) {
		y++
	}
	fmt.Println(y)
}

func isDif(y int) bool {
	y1 := y / 1000
	y2 := (y - (y1 * 1000)) / 100
	y3 := (y - (y1 * 1000) - (y2 * 100)) / 10
	y4 := y - (y1 * 1000) - (y2 * 100) - (y3 * 10)

	if y1 != y2 && y3 != y4 && y1 != y3 && y2 != y4 && y1 != y4 && y2 != y3 {
		return true
	}
	return false
}
