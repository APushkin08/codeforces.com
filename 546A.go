// https://codeforces.com/problemset/problem/546/A
package main

import (
	"fmt"
)

func main() {
	var k, n, w, buf int //k-коэффициент; n - баксы в налиции; w - желаемое количество бананана
	fmt.Scan(&k, &n, &w)

	for i := 1; i <= w; i++ {
		buf += (i * k) // подсчет долларов, для покупки необходимого числа бананов
	}

	buf -= n

	if buf > 0 {
		fmt.Println(buf)
	} else {
		fmt.Printf("%c", '0')
	}

}
