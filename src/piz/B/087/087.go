package main
import (
	."fmt"
	"strconv"
	"strings"
	"math"
)

func transpose(slice [][]string) [][]string {
    xl := len(slice[0])
    yl := len(slice)
    result := make([][]string, xl)
    for i := range result {
        result[i] = make([]string, yl)
    }
    for i := 0; i < xl; i++ {
        for j := 0; j < yl; j++ {
            result[i][j] = slice[j][i]
        }
    }
    return result
}

func main() {
	var h, w, k, maxX, maxY int
	Scan(&h, &w, &k)

	grid := make([][]string, h)
	for i := range grid {
		grid[i] = make([]string, 1)

		for j := range grid[i] {
			Scan(&grid[i][j])
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j <= w - k; j++ {
			for _, s := range grid[i] {
				n, _ := strconv.Atoi(s[j:j+k])
				if n > maxX { maxX = n }
			}
		}
	}

	tmp := make([][]string, w)
	for i := range tmp {
        tmp[i] = make([]string, 0)
    }
	for i := 0; i < h; i++ {
		for _, e := range grid[i] {
			str := strings.Split(e, "")

			for _, s := range str {
				tmp[i] = append(tmp[i], s)
			}
		}
	}
	ary := transpose(tmp)

	for i := 0; i < h; i++ {
		for j := 0; j <= w - k; j++ {
			s := ary[i][j:j+k]
			n, _ := strconv.Atoi(strings.Join(s, ""))
			if n > maxY { maxY = n }
		}
	}

	Println(math.Max(float64(maxX), float64(maxY)))
}