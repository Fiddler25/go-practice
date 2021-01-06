// https://atcoder.jp/contests/abc158/tasks/abc158_c

package main
import . "fmt"

func main() {
	var a, b int
	var i float64
	Scan(&a, &b)

	for i = 1; i <= 1000; i++ {
		if int(i * 0.08) == a && int(i * 0.1) == b {
			Println(i)
			return
		}
	}
	Println(-1)
}