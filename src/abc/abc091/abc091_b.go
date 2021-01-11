// https://atcoder.jp/contests/abc088/tasks/abc091_b

package main
import . "fmt"

func main() {
	var n, ans int
	var s string
	dic := make(map[string]int)

	Scan(&n)
	for i := 0; i < n; i++ {
		Scan(&s); dic[s]++
	}

	Scan(&n)
	for i := 0; i < n; i++ {
		Scan(&s); dic[s]--
	}

	for _, v := range dic {
		if ans < v {
			ans = v
		}
	}

	Println(ans)
}