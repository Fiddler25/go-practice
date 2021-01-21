package main
import (
	. "fmt"
	"strings"
)

func sum(ary []int) int {
	cnt := 0
	for _, e := range ary {
		cnt += e
	}
	return cnt
}

func main() {
	var h, w, cnt int
	Scan(&h, &w)

	grid := make([][]int, h)
	for i := range grid {
		grid[i] = make([]int, w)

		for j := range grid[i] {
			Scan(&grid[i][j])
		}
	}
	
	ans := make([][]string, h)
	for i := 0; i < h; i++ {
		for j := 1; j < w; j++ {
			if sum(grid[i][:j]) != sum(grid[i][j:]) { continue }

			ans[i] = make([]string, w)
			for k := 0; k < w; k++ {
				if k < j {
					ans[i][k] = "A"
				} else {
					ans[i][k] = "B"
				}
			}
			cnt ++
		}
	}

	if cnt == h {
		Println("Yes")

		for _, ary := range ans {
			Println(strings.Join(ary, ""))
		}
		return
	}

	Println("No")
}