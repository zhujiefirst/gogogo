package sort

import (
	"fmt"
	// "reflect"
)

type BaseSort struct {
	compareCount int
	exchCount    int
}

func (this *BaseSort) less(lhs Comparable, rhs Comparable, compare Compare) bool {
	this.compareCount++
	return compare(lhs, rhs) < 0
}

func (this *BaseSort) exch(a []Comparable, i int, j int) {
	this.exchCount++
	t := a[j]
	a[j] = a[i]
	a[i] = t
}

func (this *BaseSort) Show(a []Comparable) {
	for i := 0; i < len(a); i++ {
		fmt.Printf("%v ", a[i])
	}
	fmt.Println("")
}

func (this *BaseSort) IsSorted(a []Comparable, compare Compare) bool {
	for i := 0; i < len(a); i++ {
		if this.less(a[i], a[i-1], compare) == true {
			return false
		}
	}
	return true
}
