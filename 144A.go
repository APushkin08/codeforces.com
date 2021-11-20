// https://codeforces.com/problemset/problem/144/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, res int
	fmt.Scan(&n)

	arr := make([]int, n)
	max := 0
	min := 101
	leftMax := 101
	rightMin := 0

	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])

		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	for i, v := range arr {
		if v == max {
			leftMax = i
			break
		}
	}

	for i := n - 1; i >= 0; i-- {
		if arr[i] == min {
			rightMin = i
			break
		}
	}

	if leftMax > rightMin {
		res = leftMax + n - rightMin - 2
	} else {
		res = leftMax + n - rightMin - 1
	}

	fmt.Printf("%d", res)

}
