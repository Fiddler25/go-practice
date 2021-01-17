// https://atcoder.jp/contests/abc167/tasks/abc167_b

package main
import (
	. "fmt"
	"math"
)

func main() {
	var a, b, _c, k float64
	Scan(&a, &b, &_c, &k)

	min := int(math.Min(a, k))
	max := int(math.Max(0, k - a - b))

	Println(min - max)
}