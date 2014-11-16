package algrithom

import (
	"testing"
)

func TestNoRepeatRand(*testing.T) {
	N := 10000

	forceMethod(N)
	arrayMethod(N)
	listMethod(N)
	betterMethod(N)
}
