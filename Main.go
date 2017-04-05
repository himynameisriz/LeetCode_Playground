package main

import (
	"fmt"
	modules "./Modules"
)

func main() {
	fb := modules.FizzBuzz{}
	fmt.Println(fb.String(15))

	// should work, breaks on leetcode
	capsLetters := modules.AllCapsLetters{}
	fmt.Println("Detecting all caps", capsLetters.DetectCapitalUse("g"))
}