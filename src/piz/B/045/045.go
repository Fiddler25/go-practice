package main
import . "fmt"

func main() {
	var m, n int
	Scan(&m, &n)

	a, b := 0, 0
	for i := 0; i < m; i++ {
		if a + b > 99 {
			a++
			b = 0
		}

		Printf("%d + %d =\n", a, b)
		b++
	}

	a, b = 0, 0
	for i := 0; i < n; i++ {
		if a - b < 0 {
			a++
			b = 0
		}

		Printf("%d - %d =\n", a, b)
		b++
	}
}