// https://codeforces.com/problemset/problem/160/A
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	qsort(s)

	// Нахождение сумм достоинств всех монет
	var sumCoins int
	for i := 0; i < len(s); i++ {
		sumCoins += s[i]
	}

	half := float32(sumCoins) / 2
	var nCoins int
	for i := len(s) - 1; float32(sumCoins) >= half; i-- {
		sumCoins -= s[i]
		nCoins++
	}
	fmt.Println(nCoins)
}

// Быстрая сортировака
func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}
