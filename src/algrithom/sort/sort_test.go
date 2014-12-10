package sort

import (
	"math/rand"
	"testing"
)

func compareTo(lhs Comparable, rhs Comparable) int {
	a := lhs.(int)
	b := rhs.(int)
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

func genRandom(n int) []Comparable {
	var a []Comparable
	for i := 0; i < n; i++ {
		e := rand.Int() % (2 * n)
		a = append(a, e)
	}
	return a
}

func TestSort(*testing.T) {
	c := genRandom(10)

	sort := new(BaseSort)
	sort.Show(c)
}
