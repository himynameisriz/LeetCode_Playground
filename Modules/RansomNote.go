package Modules

import "fmt"

type RansomNote struct {}

func (r RansomNote) CanConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)
	for i:=0; i < len([]byte(magazine));  i++{
		char := magazine[i]
		if _, ok := m[char]; ok {
			m[char]++
		} else {
			m[char] = 1
		}
	}
	fmt.Println(m)
	for i := 0; i < len([]byte(ransomNote)); i++ {
		char := ransomNote[i]
		if val, ok := m[char]; ok && val != 0 {
			m[char]--
		} else {
			return false
		}
	}
	return true
}