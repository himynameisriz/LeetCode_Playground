package Modules

import (
	"math"
	"fmt"
	"reflect"
)

type ReverseInt struct {}

func (r ReverseInt) Reverse(x int) int {
	y := 0
	fmt.Println(reflect.TypeOf(y))
	for {
		y = y * 10 + x % 10
		x = x/10
		if y > math.MaxInt32 || y < math.MinInt32 {
			fmt.Println("Int32 overflow", y, "new typing:", reflect.TypeOf(y)) // comes out to be the same typing??
			return 0
		}
		if x == 0 {
			break
		}
	}
	return y
}