package main
import . "fmt"

func main() {
	var n, a, t, q, x, k int
	Scan(&n)

	cost := make(map[int]int)
	for i := 1; i <= n; i++ {
		Scan(&a)
		cost[i] = a
	}

	Scan(&t, &q)
	for i := 0; i < q; i++ {
		Scan(&x, &k)

		c := cost[x] * k
		if t - c < 0 { continue }
		
		t -= c
	}
	Println(t)
}