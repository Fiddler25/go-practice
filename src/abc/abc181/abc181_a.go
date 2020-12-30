// https://atcoder.jp/contests/abc181/tasks/abc181_a

package main
import . "fmt"

func main() {
	var n int
	Scan(&n)
	color := "White"
	if n % 2 == 1 {color = "Black"}

	Println(color)
}