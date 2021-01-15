// https://atcoder.jp/contests/abc176/tasks/abc176_c

package main
import . "fmt"

func main() {
	var n, ans, h int
	Scan(&n)

	ary := make([]int, n)
	for i := range ary { Scan(&ary[i]) }

	for i := 0; i < n-1; i++ {
		if ary[i] <= ary[i+1] { continue }
		h = ary[i] - ary[i+1]
		ans += h
		ary[i+1] += h
	}

	Println(ans)
}