// https://atcoder.jp/contests/abc180/tasks/abc180_c

package main
import (
	. "fmt"
	"sort"
)

func main() {
	var n int
	Scan(&n)

	var ary []int
	for i := 1; i * i <= n; i++ {
		if n % i != 0 { continue }
		ary = append(ary, i)
		
		if i * i != n {
			ary = append(ary, n / i)
		} 
	}
	sort.Ints(ary)

	for _, ans := range ary {
		Println(ans)
	}
}