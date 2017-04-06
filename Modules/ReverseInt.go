package Modules

import "math"

type ReverseInt struct {}

func (r ReverseInt) Reverse(x int) int {
	y := 0
	for {
		y = y * 10 + x % 10
		x = x/10
		if y > math.MaxInt32 || y < math.MinInt32 {
			return 0
		}
		if x == 0 {
			break
		}
	}
	return y
}