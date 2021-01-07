// https://atcoder.jp/contests/abc155/tasks/abc155_c

package main
import (
	. "fmt"
	"sort"
)

func main() {
	var n, max int
	Scan(&n)
	
	var s string
	dic := make(map[string]int)
	for i := 0; i < n; i++ {
		Scan(&s)
		dic[s]++

		if dic[s] > max { max = dic[s] }
	}

	var ary []string
	for k, v := range dic {
		if v == max { ary = append(ary, k) }
	}

	sort.Strings(ary)
	for _, ans := range ary {
		Println(ans)
	}
}