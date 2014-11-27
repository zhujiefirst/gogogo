package algrithom

import (
	"fmt"
	"testing"
)

func TestUnitFind(*testing.T) {
	N := 10
	pairs := [11]struct {
		p int
		q int
	}{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{8, 9},
		{5, 0},
		{7, 2},
		{6, 1},
		{1, 0},
		{6, 7},
	}

	fmt.Println("Unit Find -- Quick Find begin ...")
	uf := new(UFQuickFind)
	uf.Init(N)
	for _, v := range pairs {
		if uf.Connected(v.p, v.q) {
			fmt.Printf("[PASS] %v, %v\n", v.p, v.q)
			continue
		}
		uf.Unit(v.p, v.q)
		fmt.Printf("[CONN] %v, %v\n", v.p, v.q)
	}
}
