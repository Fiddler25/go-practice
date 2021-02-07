package main
import ."fmt"

func main() {
	var _a, b, n int
	Scan(&_a, &b, &n)

	pins := make([]int, n)
	for i := range pins { Scan(&pins[i]) }

	var score, rest, ans int
	firstFlg, strike, pre, spare := true, false, false, false

	for _, pin := range pins {
		score = pin
		if strike || spare { score += pin }
		if pre { score += pin }

		pre, strike, spare = strike, false, false

		if firstFlg && b == pin {
			strike = true
		} else if firstFlg {
			firstFlg = false
			rest = b - pin
		} else {
			if rest == pin { spare = true }
			firstFlg = true
		}

		ans += score
	}

	Println(ans)
}