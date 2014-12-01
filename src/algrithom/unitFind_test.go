package algrithom

import (
	"fmt"
	"testing"
	"time"
)

func TestUnitFind(*testing.T) {
	type Pairs [11]struct {
		p int
		q int
	}

	process := func(n int, pairs Pairs, uf UFI) {
		uf.Init(n)
		for _, v := range pairs {
			if uf.Connected(uf.Find, v.p, v.q) {
				fmt.Printf("[PASS] %v, %v\n", v.p, v.q)
				continue
			}
			uf.Unit(v.p, v.q)
			fmt.Printf("[CONN] %v, %v\n", v.p, v.q)
		}
	}

	N := 10
	pairs := Pairs{
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
	tb := time.Now()
	var uf UFI = new(UFQuickFind)
	process(N, pairs, uf)
	fmt.Printf("Unit Find -- Quick Find cost time: %v\n", time.Now().Sub(tb))

	fmt.Println("Unit Find -- Quick Unit begin ...")
	tb = time.Now()
	uf = new(UFQuickUnit)
	process(N, pairs, uf)
	fmt.Printf("Unit Find -- Quick Find cost time: %v\n", time.Now().Sub(tb))

	fmt.Println("Unit Find -- Weight Quick Unit begin ...")
	tb = time.Now()
	uf = new(UFWeightQuickUnit)
	process(N, pairs, uf)
	fmt.Printf("Unit Find -- Quick Find cost time: %v\n", time.Now().Sub(tb))
}
