package main
import (
	."fmt"
	"sort"
)

func include(ary []int, e int) bool {
	for _, a := range ary {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	var _n, m, k, ans int
	Scan(&_n, &m)

	ab := make([][]int, m)
	for i := 0; i < m; i++ {
		ab[i] = make([]int, 2)

		for j := 0; j < 2; j++ {
			Scan(&ab[i][j])
		}
	}

	Scan(&k)
	cd := make([][]int, k)
	for i := 0; i < k; i++ {
		cd[i] = make([]int, 2)

		for j := 0; j < 2; j++ {
			Scan(&cd[i][j])
		}
	}

	var ary []int
	for _, ele := range cd {
		for _, e := range ele {
			if !include(ary, e) {
				ary = append(ary, e)
			}
		}
	}
	sort.Ints(ary)

	for _, ele := range ab {
		cnt := 0

		for _, e := range ele {
			if include(ary[:k], e) { cnt++ }			
			if cnt == 2 { ans++ }
		}
	}

	Println(ans)
}