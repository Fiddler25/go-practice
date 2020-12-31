// https://atcoder.jp/contests/abc185/tasks/abc185_b

package main
import . "fmt"

func main() {
	var n, m, t, a, b int
	Scan(&n, &m, &t)
	r := n
	pb := 0
	for i := 0; i < m; i++ {
		Scan(&a, &b)
		if i == 0 {
			r -= a
		} else {
			r -= a - pb
		}
		if r <= 0 {break}

		pb = b
		r += b - a
		if r > n {r = n}

		if i == m - 1 {r -= t - b}
	}
	
	if r <= 0 {
		Printf("No\n")
		return
	}
	Printf("Yes\n")
}