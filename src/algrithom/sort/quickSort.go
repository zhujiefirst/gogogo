package sort

import (
// "fmt"
)

type QuickSort struct {
	BaseSort
}

func (this *QuickSort) partition(a []Comparable, compare Compare, lo int, hi int) int {
	i := lo + 1
	j := hi
	v := a[lo]
	for true {
		for this.less(a[i], v, compare) == true {
			if i == hi {
				break
			}
			i++
		}
		for this.less(a[j], v, compare) == false {
			if j == lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		this.exch(a, i, j)
		i++
		j--
	}
	this.exch(a, lo, j)

	return j
}

func (this *QuickSort) sort(a []Comparable, compare Compare, lo int, hi int) {
	if hi <= lo {
		return
	}

	p := this.partition(a, compare, lo, hi)
	this.sort(a, compare, lo, p-1)
	this.sort(a, compare, p+1, hi)
}

func (this *QuickSort) Sort(a []Comparable, compare Compare) {
	this.sort(a, compare, 0, len(a)-1)
}
