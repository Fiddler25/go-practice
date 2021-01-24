package main
import (
	."fmt"
	"strconv"
	"math"
)

func dir(s string) int {
	switch s {
	case "N":
		return 1
	case "E":
		return 2
	case "S":
		return 3
	case "W":
		return 4
	}
	return 0
}

func main() {
	var n int
	Scan(&n)

	from := make([]string, 2)
	for i := range from { Scan(&from[i]) }
	fromNum, _ := strconv.Atoi(from[0])

	to := make([]string, 2)
	for i := range to { Scan(&to[i]) }
	toNum, _ := strconv.Atoi(to[0])

	d := dir(from[1]) - dir(to[1])

	if d == 0 { // 同じ方角
		r := math.Abs(float64(fromNum - toNum)) * 100
		Println(r)
	} else if d % 2 == 0 { // 反対の方角
		Println((fromNum + toNum) * 100)
	} else {
		fromF, toF := float64(fromNum), float64(toNum)
		r := math.Abs(fromF - toF) * 100
		Println(r + (math.Min(fromF, toF) * 100 * math.Pi / 2))
	}
}