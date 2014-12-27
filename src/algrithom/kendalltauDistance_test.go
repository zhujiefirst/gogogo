package algrithom

import (
	// "log"
	"testing"
)

func TestKendalltauDistance(t *testing.T) {
	a := []int{0, 3, 1, 6, 2, 5, 4}
	b := []int{1, 0, 3, 6, 4, 2, 5}
	// log.Printf("list A: %v\nlist B: %v\n 's Kendall tau distance is %d	", a, b, KendalltauDistance(a, b))
	if KendalltauDistance(a, b) != 4 {
		t.Errorf("两序列的Kendall tau距离!=4")
	}
}
