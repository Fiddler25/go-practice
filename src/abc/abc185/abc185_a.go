// https://atcoder.jp/contests/abc185/tasks/abc185_a

package main
import . "fmt"

func main() {
	ary := make([]int, 4)
	ans := 100
	
	for i := range ary {
		Scan(&ary[i])
		if ary[i] < ans {ans = ary[i]}
	}
	Println(ans)
}