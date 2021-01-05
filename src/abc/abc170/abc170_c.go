// https://atcoder.jp/contests/abc170/tasks/abc170_c

package main
import (
	. "fmt"
	"sort"
	"math"
)

func main() {
	var x float64
	var n int
	Scan(&x, &n)

	if n == 0 {
		Println(x)
		return
	}

	ary := make([]int, n)
	for i := range ary { Scan(&ary[i]) }
	sort.Ints(ary)
	var min float64 = 100
	var ans float64 = 100

	for i := ary[0] - 1; i <= ary[len(ary) - 1] + 1; i++ {
		i := float64(i)
		flg := false
		for _, e := range ary {
			if e == int(i) {
				flg = true
				break
			}
		}
		if flg == true { continue }

		if min > math.Abs(x - i) {
			min = math.Abs(x - i)
			ans = i
		}
	}

	Println(ans)
}