package sort

import (
// "fmt"
)

type HeapSort struct {
	BaseSort
}

func (this *HeapSort) sink(a []Comparable, i int, j int, compare Compare) {
	b := a[i:j]
	b = append(make([]Comparable, 1), b...)
	size := len(b) - 1

	func(idx int) {
		for 2*idx <= size {
			child := 2 * idx
			// fmt.Println(idx, child, size)
			if child < size && compare(b[child], b[child+1]) < 0 {
				child++
			}
			if compare(b[idx], b[child]) < 0 {
				this.exch(b, idx, child)
				idx = child
				continue
			}
			break
		}
	}(1)

	copy(a[i:j], b[1:])
}

func (this *HeapSort) Sort(a []Comparable, compare Compare) {
	n := len(a)

	for i := n / 2; i >= 0; i-- {
		this.sink(a, i, n, compare)
	}

	for i := n - 1; i > 1; {
		this.exch(a, 0, i)
		i--
		this.sink(a, 0, i, compare)
	}
}
