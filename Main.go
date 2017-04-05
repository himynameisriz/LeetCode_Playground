package main

import (
	"fmt"
	modules "./Modules"
)

func main() {
	fb := modules.FizzBuzz{}
	fmt.Println(fb.String(15))
}