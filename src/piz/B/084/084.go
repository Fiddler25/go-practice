package main
import (
	. "fmt"
	"sort"
)

func main() {
	var n, m, k int
	Scan(&n, &m, &k)

	grid := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		grid[i] = make([]int, n)

		for j := 0; j < n; j++ {
			Scan(&grid[i][j])
		}
	}

	tmp := make(map[int]int)
	for i := 1; i <= m; i++ {
		cnt := 0
		var target []int

		for j := 0; j < n; j++ {
			if grid[0][j] == 3 && grid[i][j] == 3 {
				cnt++
			}

			if grid[0][j] == 0 && grid[i][j] == 3 {
				target = append(target, j+1)
			}
		}

		if cnt >= k {
			for _, e := range target {
				tmp[e] = 0
			}
		}
	}

	var ans []int
	for e := range tmp {
		ans = append(ans, e)
	}
	sort.Ints(ans)

	if len(ans) > 0 {
		a := Sprint(ans)
		Println(a[1:len(a)-1])
	} else {
		Println("no")
	}
}