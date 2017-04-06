package main

import (
	"fmt"
	modules "./Modules"
)

func main() {
	//fb := modules.FizzBuzz{}
	//fmt.Println(fb.String(15))
	//
	//// should work, breaks on leetcode
	//capsLetters := modules.AllCapsLetters{}
	//fmt.Println("Detecting all caps", capsLetters.DetectCapitalUse("g"))
	//
	//cwn := modules.CanWinNim{}
	//fmt.Println("Can win nim?", cwn.CanWinNim(4))

	//ransomNote := modules.RansomNote{}
	//fmt.Println("Can the ransom note be completed?", ransomNote.CanConstruct("fffbfg", "effjfggbffjdgbjjhhdegh"))

	reverseInt := modules.ReverseInt{}
	fmt.Println("Reverse int: ", reverseInt.Reverse(234))

}
