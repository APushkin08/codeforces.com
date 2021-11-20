// https://codeforces.com/problemset/problem/510/A
package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	mtx := make([][]rune, n)
	for i := range mtx {
		mtx[i] = make([]rune, m)
	}

	for i := range mtx {
		for j := 0; j < m; j++ {
			mtx[i][j] = '#'
		}
	}

	for i := 1; i < n-1; i += 2 {
		if (i+1)%4 == 0 {
			for j := 1; j < m; j++ {
				mtx[i][j] = '.'
			}
		} else {
			for j := 0; j < m-1; j++ {
				mtx[i][j] = '.'
			}
		}
	}

	for i := range mtx {
		for j := 0; j < m; j++ {
			fmt.Printf("%c", mtx[i][j])
		}
		fmt.Println()
	}
}
