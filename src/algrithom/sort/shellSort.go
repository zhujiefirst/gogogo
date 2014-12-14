package sort

import (
// "fmt"
)

type ShellSort struct {
	BaseSort
}

func (this *ShellSort) Sort(a []Comparable, compare Compare) {
	arrayLen := len(a)

	h := 1
	for h < arrayLen/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := 1; i < len(a); i++ {
			for j := i; j >= h && this.less(a[j], a[j-h], compare); j -= h {
				this.exch(a, j, j-h)
			}
		}
		h = h / 3
	}
}
