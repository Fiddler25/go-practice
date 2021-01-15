package main
import (
	. "fmt"
	"sort"
)

func ul(y, x int, grid []string) int {
	cnt := 0
	area := y * x

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if string(grid[i][j]) == "@" { cnt++ }
		}
	}
	return area + cnt * cnt
}

func ur(y, x, mx int, grid []string) int {
	cnt := 0
	area := y * (mx - x)

	for i := 0; i < y; i++ {
		for j := x; j < mx; j++ {
			if string(grid[i][j]) == "@" { cnt++ }
		}
	}
	return area + cnt * cnt
}

func ll(y, x, my int, grid []string) int {
	cnt := 0
	area := (my - y) * x

	for i := y; i < my; i++ {
		for j := 0; j < x; j++ {
			if string(grid[i][j]) == "@" { cnt++ }
		}
	}
	return area + cnt * cnt
}

func lr(y, x, my, mx int, grid []string) int {
	cnt := 0
	area := (my - y) * (mx - x)

	for i := y; i < my; i++ {
		for j := x; j < mx; j++ {
			if string(grid[i][j]) == "@" { cnt++ }
		}
	}
	return area + cnt * cnt
}

func main() {
	var h, w int
	Scan(&h, &w)

	grid := make([]string, h)
	for s := range grid { Scan(&grid[s]) }

	var a, b int
	min := 10000000000
	for y := 1; y < h; y++ {
		for x := 1; x < w; x++ {
			ary := []int {
				ul(y, x, grid),
				ur(y, x, w, grid),
				ll(y, x, h, grid),
				lr(y, x, h, w, grid),
			}
			sort.Ints(ary)

			m := ary[3] - ary[0]
			if m < min {
				a = y
				b = x
				min = m
			}
		}
	}

	Println(a, b)
}