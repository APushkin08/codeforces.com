// https://codeforces.com/problemset/problem/1154/A
package main

import (
	"fmt"
)

func main() {
	var mas [4]uint32

	for k := 0; k < len(mas); k++ {
		fmt.Scan(&mas[k])
	}
	//+сортировка пузырьком
	var i = 0
	var swapCnt = false

	for i < len(mas) {
		if i+1 != len(mas) && mas[i] > mas[i+1] {
			mas[i], mas[i+1] = mas[i+1], mas[i]
			swapCnt = true
		}
		i++
		if i == len(mas) && swapCnt == true {
			swapCnt = false
			i = 0
		}
	}
	//-сортировка пузырьком

	fmt.Println(mas[3]-mas[0], mas[3]-mas[1], mas[3]-mas[2])

}
