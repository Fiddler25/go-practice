// https://atcoder.jp/contests/abc156/tasks/abc156_c

package main
import . "fmt"

func main() {
	var n int
	Scan(&n)

	ary := make([]int, n)
	for e := range ary { Scan(&ary[e]) }

	ans := 100000000
	for p := 1; p <= 100; p++ {
		sum := 0

		for _, x := range ary {
			sum += (x - p) * (x - p)
		}

		if sum < ans { ans = sum }
	}

	Println(ans)
}