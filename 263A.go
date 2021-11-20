// https://codeforces.com/problemset/problem/263/A
package main

import (
	"fmt"
	"math"
)

func main() {
	var matrix [5][5]int
	var k, l int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&matrix[i][j])
			if matrix[i][j] == 1 {
				k = i - 2
				l = j - 2
			}
		}
	}

	rez := math.Abs(float64(k)) + math.Abs(float64(l))
	fmt.Println(rez)
}
