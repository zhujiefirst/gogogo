package sort

import (
// "fmt"
)

type MergeSort struct {
	BaseSort
}

func (this *MergeSort) merge(a []Comparable, compare Compare, lo int, mid int, hi int) {
	i := lo
	j := mid + 1

	arrayBak := make([]Comparable, len(a))
	copy(arrayBak, a)

	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = arrayBak[j]
			j++
		} else if j > hi {
			a[k] = arrayBak[i]
			i++
		} else if this.less(arrayBak[i], arrayBak[j], compare) {
			a[k] = arrayBak[i]
			i++
		} else {
			a[k] = arrayBak[j]
			j++
		}
	}
}

func (this *MergeSort) sort(a []Comparable, compare Compare, lo int, hi int) {
	if lo >= hi {
		return
	}

	mid := (lo + hi) / 2

	this.sort(a, compare, lo, mid)
	this.sort(a, compare, mid+1, hi)
	this.merge(a, compare, lo, mid, hi)
}

func (this *MergeSort) Sort(a []Comparable, compare Compare) {
	this.sort(a, compare, 0, len(a)-1)
}
