package priority_queue

import (
	"fmt"
)

type HeapPQ struct {
	vals    []Comparable
	compare Compare
}

func (this *HeapPQ) Init(keys []Comparable, comp Compare) {
	this.vals = make([]Comparable, 1)
	if keys != nil {
		for i := 0; i < len(keys); i++ {
			this.Insert(keys[i])
		}
	}

	this.compare = comp
}

func (this *HeapPQ) Insert(key Comparable) {
	this.vals = append(this.vals, key)
	this.swim(this.Size())

	fmt.Println(this.vals)
}

func (this *HeapPQ) Max() Comparable {
	if this.IsEmpty() {
		return nil
	}

	return this.vals[1]
}

func (this *HeapPQ) DelMax() Comparable {
	if this.IsEmpty() {
		return nil
	}

	max := this.Max()
	this.exch(1, this.Size())
	this.vals = this.vals[:len(this.vals)-1]
	this.sink(1)

	return max
}

func (this *HeapPQ) IsEmpty() bool {
	return this.Size() == 0
}

func (this *HeapPQ) Size() int {
	return len(this.vals) - 1
}

func (this *HeapPQ) less(lhsIdx int, rhsIdx int) bool {
	return this.compare(this.vals[lhsIdx], this.vals[rhsIdx]) < 0
}

func (this *HeapPQ) exch(lhsIdx int, rhsIdx int) {
	if lhsIdx <= 0 || rhsIdx <= 0 || lhsIdx == rhsIdx {
		return
	}

	tmp := this.vals[lhsIdx]
	this.vals[lhsIdx] = this.vals[rhsIdx]
	this.vals[rhsIdx] = tmp
}

func (this *HeapPQ) swim(idx int) {
	for idx > 1 && this.less(idx/2, idx) == true {
		this.exch(idx/2, idx)
		idx /= 2
	}
}

func (this *HeapPQ) sink(idx int) {
	for 2*idx <= this.Size() {
		child := 2 * idx
		if child < this.Size() && this.less(child, child+1) == true {
			child++
		}
		if this.less(idx, child) != true {
			break
		}
		this.exch(idx, child)
		idx = child
	}
}
