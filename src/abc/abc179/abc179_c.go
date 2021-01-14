// https://atcoder.jp/contests/abc179/tasks/abc179_c

package main
import . "fmt"

func main() {
	var n, ans int
	Scan(&n)

	for a := 1; a <= n; a++ {
		ans += (n - 1) / a
	}

	Println(ans)
}