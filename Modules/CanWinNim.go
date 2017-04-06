package Modules

type CanWinNim struct {}

func (c CanWinNim) CanWinNim(n int) bool {
	return n % 4 != 0
}