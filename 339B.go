// https://codeforces.com/problemset/problem/339/B
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, next int // количество домов, количество дел, следующее дело
	var rez int64      // количество шагов
	crnt := 1          // ксюха в 1м доме
	fmt.Scan(&n, &m)
	// проходимся по всем делам и сразу же вычисляем длинну маршрута
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < m; i++ {
		//fmt.Scan(&next) // считываем адрес следующего дела
		fmt.Fscan(in, &next)
		if crnt < next {
			rez += int64(next - crnt) // добавляем в результат разницу меж тек. положением и сл.
			crnt = next               // теперь катюха в сл. доме
		} else if crnt > next {
			rez += int64(n - (crnt - next))
			crnt = next
		}
		// если следующее дело находится в текущем доме, то шаги не увеличиваются
	}
	fmt.Println(rez)
}
