// https://codeforces.com/problemset/problem/443/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')

	mp := make(map[byte]int)

	for i := 1; i < len(line)-2; i += 3 {
		mp[line[i]]++
	}
	fmt.Println(len(mp))
}
