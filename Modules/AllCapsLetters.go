package Modules

import "regexp"

type AllCapsLetters struct {}

func (a AllCapsLetters) DetectCapitalUse(word string) bool {
	matched, _ := regexp.Match("^[^a-z]*$", []byte(word))
	return matched
}

// ^[^a-z]*$ <- our regex