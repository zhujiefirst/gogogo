package priority_queue

import (
	"fmt"
)

type ArrayPQ struct {
	vals    []Comparable
	compare Compare
}

func (this *ArrayPQ) Init(keys []Comparable, comp Compare) {
	this.vals = make([]Comparable, 0)

	if keys == nil {
		for i := 0; i < len(keys); i++ {
			this.Insert(keys[i])
		}
	}

	this.compare = comp
}

func (this *ArrayPQ) Insert(key Comparable) {
	this.vals = append(this.vals, key)
	for i := this.Size() - 1; i > 0; i-- {
		if this.less(this.vals[i], this.vals[i-1]) {
			this.exch(i, i-1)
		}
	}
	fmt.Println(this.vals)
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
	if lhsIdx < 0 || rhsIdx < 0 || lhsIdx == rhsIdx {
		return
	}

	tmp := this.vals[lhsIdx]
	this.vals[lhsIdx] = this.vals[rhsIdx]
	this.vals[rhsIdx] = tmp
}
