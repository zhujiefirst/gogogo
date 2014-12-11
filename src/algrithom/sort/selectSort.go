package sort

import (
// "fmt"
)

type SelectSort struct {
	BaseSort
}

func (this *SelectSort) Sort(a []Comparable, compare Compare) {
	for i := 0; i < len(a); i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if this.less(a[j], a[min], compare) == true {
				min = j
			}
		}
		this.exch(a, i, min)
	}
}
