// https://codeforces.com/problemset/problem/718/C
package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {

	var a = []*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1), big.NewInt(0)}
	var v = []*big.Int{big.NewInt(1), big.NewInt(0)}

	var n, m, tp, l, r int
	var x, buf int64
	var toOut *big.Int

	fmt.Scan(&n, &m)

	arr := make([]*big.Int, n)
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &buf)
		arr[i] = big.NewInt(buf)
	}

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &tp)
		if tp == 1 {
			fmt.Fscan(in, &l, &r, &x)
			for j := l; j <= r; j++ {
				arr[j-1].Add(arr[j-1], big.NewInt(x))
			}

		} else {
			fmt.Fscan(in, &l, &r)
			toOut = big.NewInt(0)
			for j := l; j <= r; j++ {
				if arr[j-1].Cmp(big.NewInt(1)) == 0 {
					toOut.Add(toOut, big.NewInt(1))
				} else {
					var p big.Int                  // кладем в p копию arr[j-1], во избежание изменения элемента
					p.Add(big.NewInt(0), arr[j-1]) // массива при использовании метода .Add из пакета big

					toOut.Add(toOut, mult(mtxPow(a, p.Add(&p, big.NewInt(-1))), v)[0])
				}
			}

			fmt.Println(toOut.Rem(toOut, big.NewInt(1000000007)))
		}
	}

}

func mult(x, y []*big.Int) []*big.Int {

	var res []*big.Int
	var a, b, c, d, e, f big.Int

	if len(y) == 2 {
		a.Add(e.Mul(x[0], y[0]), f.Mul(x[1], y[1]))
		b.Add(e.Mul(x[2], y[0]), f.Mul(x[3], y[1]))
		return append(res, &a, &b)
	}

	a.Add(e.Mul(x[0], y[0]), f.Mul(x[1], y[2]))
	b.Add(e.Mul(x[0], y[1]), f.Mul(x[1], y[3]))
	c.Add(e.Mul(x[2], y[0]), f.Mul(x[3], y[2]))
	d.Add(e.Mul(x[2], y[1]), f.Mul(x[3], y[3]))
	return append(res, &a, &b, &c, &d)
}

func mtxPow(x []*big.Int, n *big.Int) []*big.Int {

	var p big.Int           // изащренный способ создания копии (n), для избежания изменения его
	p.Add(big.NewInt(0), n) // значения в методe .Rem - остаток от деления из пакета big

	if n.Cmp(big.NewInt(1)) == 0 {
		return x
	}
	if p.Rem(&p, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return mtxPow(mult(x, x), n.Div(n, big.NewInt(2)))
	}
	return mult(x, mtxPow(mult(x, x), n.Div(n, big.NewInt(2))))
}
