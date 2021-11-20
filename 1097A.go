// https://codeforces.com/problemset/problem/1097/A
package main

import (
	"fmt"
)

func main() {
	var naStole string
	var naRukah [5]string
	var flag = false

	fmt.Scan(&naStole)
	for k := 0; k < len(naRukah); k++ {
		fmt.Scan(&naRukah[k])
	}

	for _, val := range naRukah {
		if string([]rune(val)[0]) == string([]rune(naStole)[0]) {
			flag = true
			break
		}
		if string([]rune(val)[1]) == string([]rune(naStole)[1]) {
			flag = true
			break
		}
	}
	if flag == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
	// //st := naRukah[k][1]
	// fmt.Println(string([]rune(naStole)[0]))
	// fmt.Println(naRukah)
}
