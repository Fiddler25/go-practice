// https://atcoder.jp/contests/abc179/tasks/abc179_b

package main
import . "fmt"

func main() {
	var n, a, b, cnt int
	Scan(&n)

	for i := 0; i < n; i++ {
		Scan(&a, &b)
		if a == b {
			cnt++
			continue
		}
		if cnt == 3 { break }
		
		cnt = 0
	}

	if cnt >= 3 {
		Println("Yes")
		return
	}
	Println("No")
}