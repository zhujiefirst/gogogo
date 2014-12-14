package sort

import (
// "fmt"
)

type InsertionSort struct {
	BaseSort
}

func (this *InsertionSort) Sort(a []Comparable, compare Compare) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && this.less(a[j], a[j-1], compare); j-- {
			this.exch(a, j, j-1)
		}
	}
}
