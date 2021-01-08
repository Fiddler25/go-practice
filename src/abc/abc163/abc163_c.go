// https://atcoder.jp/contests/abc163/tasks/abc163_c

package main
import . "fmt"

func main() {
	var n, a int
	Scan(&n)

	ma := make(map[int]int)
	for i := 1; i < n; i++ {
		Scan(&a)
		ma[a]++
	}
	
	for i := 1; i <= n; i++ { Println(ma[i]) }
}