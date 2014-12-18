package priority_queue

import (
	"fmt"
	"testing"
)

func TestPQ(*testing.T) {
	compare := func(lhs Comparable, rhs Comparable) int {
		a := lhs.(int)
		b := rhs.(int)
		// fmt.Printf("a=%v\tb=%v\n", a, b)
		if a < b {
			return -1
		} else if a > b {
			return 1
		} else {
			return 0
		}
	}

	arrayPQAl := func() {
		arrayPQ := new(ArrayPQ)
		arrayPQ.Init(nil, compare)
		arrayPQ.Insert(10)
		arrayPQ.Insert(1)
		arrayPQ.Insert(13)
		arrayPQ.Insert(3)
		arrayPQ.Insert(2)
		fmt.Printf("Max value expect=13, actul=%v\n", arrayPQ.DelMax())
		fmt.Printf("Max value expect=10, actul=%v\n", arrayPQ.Max())
		fmt.Printf("PQ's size expect=4, actul=%v\n", arrayPQ.Size())
	}

	heapPQAl := func() {
		heapPQ := new(HeapPQ)
		heapPQ.Init(nil, compare)
		heapPQ.Insert(10)
		heapPQ.Insert(1)
		heapPQ.Insert(13)
		heapPQ.Insert(3)
		heapPQ.Insert(2)
		fmt.Printf("Max value expect=13, actul=%v\n", heapPQ.DelMax())
		fmt.Printf("Max value expect=10, actul=%v\n", heapPQ.Max())
		fmt.Printf("PQ's size expect=4, actul=%v\n", heapPQ.Size())
	}

	arrayPQAl()
	heapPQAl()
}
