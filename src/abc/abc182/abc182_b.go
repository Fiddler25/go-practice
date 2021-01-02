// https://atcoder.jp/contests/abc182/tasks/abc182_b

package main
import (
	. "fmt"
	"sort"
)

func main() {
	var n, max, ans int
	Scan(&n)

	ary := make([]int, n)
	for i := range ary { Scan(&ary[i]) }
	sort.Ints(ary)

	for i := 2; i <= ary[n-1]; i++ {
		cnt := 0
		for _, e := range ary { if e % i == 0 { cnt++ } }

		if max < cnt { 
			max = cnt
			ans = i
		}
	}

	Println(ans)
}