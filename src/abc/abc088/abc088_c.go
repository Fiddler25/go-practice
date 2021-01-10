// https://atcoder.jp/contests/abc088/tasks/abc088_c

package main
import . "fmt"

func main() {
	grid := [3][3]int{}
	for i := range grid {
		for j := range grid {
			Scan(&grid[i][j])
		}
	}

	row1 := make([]int, 3)
	for i, g := range grid[0] { row1[i] = g - 0 }

	col1 := make([]int, 3)
	for i, ar := range grid { col1[i] = ar[0] - row1[0] }

	tmp := [3][3]int{}
	for y, c := range col1 {
		for x, r := range row1 {
			tmp[y][x] = c + r
		}
	}

	flg := true
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			if grid[y][x] != tmp[y][x] { flg = false }
		}
	}

	if flg {
		Println("Yes")
		return
	}
	Println("No")
}