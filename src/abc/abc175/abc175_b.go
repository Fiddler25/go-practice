// https://atcoder.jp/contests/abc175/tasks/abc175_b

package main
import (
	. "fmt"
	"sort"
)

func main() {
	var n, ans int
	Scan(&n)
	ary := make([]int, n)
	for i := range ary { Scan(&ary[i]) }
	sort.Ints(ary)
	
	for i := 0; i < n; i++ {
		a := ary[i]

		for j := i+1; j < n; j++ {
			b := ary[j]

			for k := j+1; k < n; k++ {
				c := ary[k]
				
				if a != b && b != c && c != a {
					if a + b > c { ans++ }
				}
			}
		}
	}

	Println(ans)
}