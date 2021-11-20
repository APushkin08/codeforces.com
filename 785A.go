// https://codeforces.com/problemset/problem/785/A
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fig := map[string]int{
		"Tetrahedron":  4,
		"Cube":         6,
		"Octahedron":   8,
		"Dodecahedron": 12,
		"Icosahedron":  20,
	}

	var st string
	var i, n, rez int

	in := bufio.NewReader(os.Stdin)

	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Fscan(in, &st)
		rez += fig[st]
	}

	fmt.Println(rez)
}
