// https://atcoder.jp/contests/abc172/tasks/abc172_c

package main
import (
	. "fmt"
	"bufio"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, m, k, ans int
	Fscan(r, &n, &m, &k)

	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(r, &a[i])
		a[i] += a[i-1]
	}
	b := make([]int, m+1)
	for i := 1; i <= m; i++ {
		Fscan(r, &b[i])
		b[i] += b[i-1]
	}

	j := m
	for i := 0; i <= n; i++ {
		if a[i] > k { break }
		for a[i] + b[j] > k { j-- }

		if ans < i + j {
			ans = i + j
		}
	}

	Println(ans)
}