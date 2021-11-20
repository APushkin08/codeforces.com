// https://codeforces.com/problemset/problem/71/A
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]string, n)
	fmt.Scanln(&s[0]) //считываем первую(нулевую) строку, непонятно почему в цикле ниже не считвыается нулевой элемент слайса, надо бы разобратсья
	for i := 0; i < n; i++ {
		fmt.Scanln(&s[i])
	}
	for i := 0; i < n; i++ {
		var l int
		l = len(s[i]) //длинна текущего слова
		if len(s[i]) > 10 {
			fmt.Printf("%c%d%c\n", s[i][0], l-2, s[i][l-1])
		} else {
			fmt.Println(s[i])
		}

	}

}
