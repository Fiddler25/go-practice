// https://atcoder.jp/contests/abc186/tasks/abc186_b

package main
import . "fmt"

func main() {
	var h, w, a int
	Scan(&h, &w)
	sum := 0
	min := 101
	
	for i := 0; i < h * w; i++ {
		Scan(&a)
		sum += a
		if a < min {min = a}
	}

	Println(sum - min * h * w)
}