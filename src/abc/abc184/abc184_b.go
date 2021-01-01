// https://atcoder.jp/contests/abc184/tasks/abc184_b

package main
import . "fmt"

func main() {
	var n, x int
	var str string
	Scan(&n, &x, &str)
	
	for _, s := range str {
		if string(s) == "o" {
			x++
		} else if x > 0 {
			x--
		}
	}
	Println(x)
}