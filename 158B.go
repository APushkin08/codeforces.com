// https://codeforces.com/problemset/problem/158/B
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cnt := map[int]int{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
	}
	// + считывание и подсчет количеств групп
	var n int
	fmt.Scan(&n)
	inBuf := 0
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &inBuf)
		cnt[inBuf]++
	}
	// - считывание и подсчет количеств групп

	cntTxs := cnt[4] // количество необходимого такси
	// группируем 3ки и 1ки
	cntTxs += cnt[3]

	cnt[1] = cnt[1] - cnt[3]
	if cnt[1] < 0 {
		cnt[1] = 0
	}

	// группируем 2ки
	cntTxs += cnt[2] / 2
	if (cnt[2] % 2) != 0 { // если кол. двоек нечетное
		if (cnt[1] >= 0) && (cnt[1] <= 2) {
			cntTxs++
			cnt[1] = 0
		} else if cnt[1] > 2 {
			cntTxs++
			cnt[1] -= 2
		}
	}

	// группируем 1ки
	cntTxs += cnt[1] / 4

	if (cnt[1] % 4) > 0 {
		cntTxs++
	}

	fmt.Println(cntTxs)
}
