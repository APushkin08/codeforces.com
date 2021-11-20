// https://codeforces.com/problemset/problem/1220/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var st string
	var one, zero int
	fmt.Scan(&n)
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &st)
	for i := 0; i < n; i++ {
		if string([]rune(st)[i]) == "n" {
			one++
			continue
		}
		if string([]rune(st)[i]) == "z" {
			zero++
		}
	}
	st = ""
	for i := 0; i < one; i++ {
		st += "1 "
	}
	for i := 0; i < zero; i++ {
		st += "0 "
	}
	fmt.Print(st)
	// writer := bufio.NewWriter(os.Stdout)
	// for i := 0; i < one; i++ {
	// 	writer.WriteRune('1')
	// 	writer.WriteRune(' ')
	// }
	// for i := 0; i < zero; i++ {
	// 	writer.WriteRune('0')
	// 	writer.WriteRune(' ')
	// }
	// writer.Flush()
}
