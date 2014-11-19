package algrithom

import (
	"testing"
)

func TestNoRepeatRand(*testing.T) {
	N := 10000

	ForceMethod(N)
	ArrayMethod(N)
	ListMethod(N)
	BetterMethod(N)
}
