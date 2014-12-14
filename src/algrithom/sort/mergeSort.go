package sort

import (
	// "fmt"
	"math"
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

func (this *MergeSort) sort1(a []Comparable, compare Compare, lo int, hi int) {
	if lo >= hi {
		return
	}

	mid := (lo + hi) / 2

	this.sort1(a, compare, lo, mid)
	this.sort1(a, compare, mid+1, hi)
	this.merge(a, compare, lo, mid, hi)
}

func (this *MergeSort) sort2(a []Comparable, compare Compare) {
	for sz := 1; sz < len(a); sz += sz {
		for lo := 0; lo < len(a)-sz; lo += sz + sz {
			this.merge(a, compare, lo, lo+sz-1, int(math.Min(float64(lo+sz+sz-1), float64(len(a)-1))))
		}
	}
}

func (this *MergeSort) Sort(a []Comparable, compare Compare) {
	// this.sort1(a, compare, 0, len(a)-1)
	this.sort2(a, compare)
}
