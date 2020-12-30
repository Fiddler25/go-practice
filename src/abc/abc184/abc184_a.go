// https://atcoder.jp/contests/abc184/tasks/abc184_a

package main
import . "fmt"

func main() {
	var a, b, c, d int
	Scan(&a, &b, &c, &d)

	Println(a * d - b * c)
}