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

func (this *QuickSort) sort1(a []Comparable, compare Compare, lo int, hi int) {
	if hi <= lo {
		return
	}

	p := this.partition(a, compare, lo, hi)
	this.sort1(a, compare, lo, p-1)
	this.sort1(a, compare, p+1, hi)
}

func (this *QuickSort) sort2(a []Comparable, compare Compare, lo int, hi int) {
	if hi <= lo {
		return
	}

	if hi <= lo+5 {
		insertionSort := new(InsertionSort)
		insertionSort.Sort(a[lo:hi], compare)
	}

	p := this.partition(a, compare, lo, hi)
	this.sort2(a, compare, lo, p-1)
	this.sort2(a, compare, p+1, hi)
}

func (this *QuickSort) sort3(a []Comparable, compare Compare, lo int, hi int) {
	if hi <= lo {
		return
	}

	lt := lo
	i := lo + 1
	gt := hi

	v := a[lo]
	for i <= gt {
		cmp := compare(v, a[i])
		if cmp < 0 {
			this.exch(a, i, gt)
			gt--
		} else if cmp > 0 {
			this.exch(a, i, lt)
			i++
			lt++
		} else {
			i++
		}
	}
	this.sort3(a, compare, lo, lt-1)
	this.sort3(a, compare, gt+1, hi)
}

func (this *QuickSort) Sort(a []Comparable, compare Compare, i int) {
	if i == 1 {
		this.sort1(a, compare, 0, len(a)-1)
	} else if i == 2 {
		this.sort2(a, compare, 0, len(a)-1)
	} else {
		this.sort3(a, compare, 0, len(a)-1)
	}
}
