// https://codeforces.com/problemset/problem/339/A
package main

import (
	"fmt"
	s "strings"
)

func main() {
	var line, rez string

	var cnt1, cnt2, cnt3 int

	fmt.Scan(&line)
	/*	for i := 0; i < len(line); i += 2 {
			if string([]rune(line)[i]) == "1" {
				cnt1++
				continue
			} else if string([]rune(line)[i]) == "2" {
				cnt2++
				continue
			} else {
				cnt3++
			}
		}
	*/
	cnt1 = s.Count(line, "1")
	cnt2 = s.Count(line, "2")
	cnt3 = s.Count(line, "3")

	for i := 0; i < cnt1; i++ {
		rez += "1+"
	}
	for i := 0; i < cnt2; i++ {
		rez += "2+"
	}
	for i := 0; i < cnt3; i++ {
		rez += "3+"
	}

	rez = rez[:len(rez)-1]
	fmt.Println(rez)
}
