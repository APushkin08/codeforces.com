// https://codeforces.com/problemset/problem/118/A
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	str = strings.ToLower(str) // делаем все символы строчными

	repDel := func(r rune) rune {
		switch {
		case r == 'a' || r == 'o' || r == 'y' || r == 'e' || r == 'u' || r == 'i':
			return -1 //* удаление символа, т.е. если в strings.Map попадется один из вышеперечисленных символов, он не впишется в str
		}
		return r
	}
	str = strings.Map(repDel, str)

	for i := range str {
		fmt.Printf(".%c", str[i])
	}

}
