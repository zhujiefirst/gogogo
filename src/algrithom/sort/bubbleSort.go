package sort

import (
// "fmt"
)

type BubbleSort struct {
	BaseSort
}

func (this *BubbleSort) Sort(a []Comparable, compare Compare) {
	for i := 0; i < len(a); i++ {
		for j := 1; j < len(a)-i; j++ {
			if this.less(a[j], a[j-1], compare) {
				this.exch(a, j-1, j)
			}
		}
	}
}
