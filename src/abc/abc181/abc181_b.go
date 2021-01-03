// https://atcoder.jp/contests/abc181/tasks/abc181_b

package main
import . "fmt"

func main() {
	var n, a, b, ans int
	Scan(&n)
	
	for i := 0; i < n; i++ {
		Scan(&a, &b)
		ans += b * (b + 1) / 2 - a * (a - 1) / 2
	}

	Println(ans)
}