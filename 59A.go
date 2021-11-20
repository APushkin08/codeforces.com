// https://codeforces.com/problemset/problem/59/A
package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)

	cntLowerCase := 0

	for _, v := range input {
		if v >= 97 && v <= 122 {
			cntLowerCase++
		}
	}

	if len(input)-cntLowerCase <= cntLowerCase {
		fmt.Println(strings.ToLower(input))
	} else {
		fmt.Println(strings.ToUpper(input))
	}
}
