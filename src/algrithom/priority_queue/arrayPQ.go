package priority_queue

import (
// "fmt"
)

type ArrayPQ struct {
	vals    []Comparable
	compare Compare
}

func (this *ArrayPQ) Init(keys []Comparable, comp Compare) {
	if keys == nil {
		this.vals = make([]Comparable, 0)
	} else {
		this.vals = make([]Comparable, len(keys))
		copy(this.vals, keys)
	}

	this.compare = comp
}

func (this *ArrayPQ) Insert(key Comparable) {
	this.vals = append(this.vals, key)
	var maxIdx = this.Size() - 1
	for i := 0; i < this.Size()-1; i++ {
		if this.less(this.vals[maxIdx], this.vals[i]) {
			maxIdx = i
		}
	}
	this.exch(maxIdx, this.Size()-1)
}

func (this *ArrayPQ) Max() Comparable {
	if this.IsEmpty() {
		return nil
	}

	return this.vals[this.Size()-1]
}

func (this *ArrayPQ) DelMax() Comparable {
	if this.IsEmpty() {
		return nil
	}

	max := this.Max()
	this.vals = this.vals[:this.Size()-1]

	return max
}

func (this *ArrayPQ) IsEmpty() bool {
	return len(this.vals) == 0
}

func (this *ArrayPQ) Size() int {
	return len(this.vals)
}

func (this *ArrayPQ) less(lhs Comparable, rhs Comparable) bool {
	return this.compare(lhs, rhs) < 0
}

func (this *ArrayPQ) exch(lhsIdx int, rhsIdx int) {
	if lhsIdx == rhsIdx {
		return
	}

	tmp := this.vals[lhsIdx]
	this.vals[lhsIdx] = this.vals[rhsIdx]
	this.vals[rhsIdx] = tmp
}
