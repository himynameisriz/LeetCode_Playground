package Modules

import "strconv"

type FizzBuzz struct {

}

func getFizzOrBuss(n int) string {
	if n % 3 == 0 && n % 5 == 0 {
		return "FizzBuzz"
	} else if n % 5 == 0 {
		return "Buzz"
	} else if n % 3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}

func (f FizzBuzz) String(n int) []string {
	arr := make([]string, n)
	for i := 1; i <= n; i++ {
		arr[i-1] = getFizzOrBuss(i)
	}
	return arr
}