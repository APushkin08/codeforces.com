// https://codeforces.com/problemset/problem/236/A
package main

import (
	"fmt"
)

func main() {
	var st string
	fmt.Scan(&st)
	simb := make(map[string]int)

	for i := 0; i < len(st); i++ {
		//fmt.Print(string([]rune(st)[i]))
		simb[string([]rune(st)[i])]++
	}

	if len(simb)%2 == 0 {
		fmt.Print("CHAT WITH HER!")
	} else {
		fmt.Print("IGNORE HIM!")
	}

	// for si, am := range simb {
	// 	fmt.Printf("%s\t%d\n", si, am)
	// }

}
