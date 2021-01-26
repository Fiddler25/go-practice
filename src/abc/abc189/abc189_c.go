package main
import (
	."fmt"
	"math"
)

func main() {
	var n int
	Scan(&n)

	ary := make([]int, n)
	for i := range ary { Scan(&ary[i]) }

	var ans int
	for l := 0; l < n; l++ {
		x := ary[l]

		for r := l; r < n; r++ {
			x = int(math.Min(float64(x), float64(ary[r])))
			ans = int(math.Max(float64(ans), float64(x * (r - l + 1))))
		}
	}
	Println(ans)
}