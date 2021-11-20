// https://codeforces.com/problemset/problem/1146/A
package main

import "fmt"

func main() {
	var st string
	var aNum, notA int

	fmt.Scan(&st)

	for i := 0; i < len(st); i++ {
		if string([]rune(st)[i]) == "a" {
			aNum++
		}
	}
	notA = len(st) - aNum
	//fmt.Println(notA)
	if aNum > notA {
		fmt.Println(len(st))
	} else {
		fmt.Println(len(st) - (notA - aNum) - 1)
	}

	// fmt.Println(len(st))
	// fmt.Println(aNum)

	//fmt.Println(string([]rune(st)[len(st)-1]))

}
