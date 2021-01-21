package main
import (
	. "fmt"
	"strings"
)

func uniq(slice []int) []int {
	m := make(map[int]int)
	uniq := make([]int, 0)
	for _, e := range slice {
		if _, ok := m[e]; !ok {
			m[e] = 0
			uniq = append(uniq, e)
		}
	}
	return uniq
}

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

	var target []int
	for i := 0; i <= m; i++ {
		if i == 0 { continue }
		cnt := 0

		for j := 0; j < n; j++ {
			if grid[0][j] == 3 && grid[i][j] == 3 {
				cnt++
			}
		}

		if cnt >= k {
			target = append(target, i)
		}
	}

	var ans []int
	for _, i := range target {
		for j := 0; j < n; j++ {
			if grid[0][j] == 0 && grid[i][j] == 3 {
				ans = append(ans, j+1)
			}
		}
	}

	if len(ans) == 0 {
		Println("no")
		return
	}

	ans = uniq(ans)

	str := ""
    for _, a := range ans {
        str += Sprintf("%d ", a)
	}
	str = strings.TrimRight(str, " ")

	Println(str)
}