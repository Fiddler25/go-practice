// https://atcoder.jp/contests/abc188/tasks/abc188_b

package main
import . "fmt"

func main() {
	var n, ans int
	Scan(&n)

	a := make([]int, n)
	for i := range a { Scan(&a[i]) }

	b := make([]int, n)
	for i := range b { Scan(&b[i]) }

	for i := 0; i < n; i++ {
		ans += a[i] * b[i]
	}
	
	if ans == 0 {
		Println("Yes")
		return
	}
	Println("No")
}