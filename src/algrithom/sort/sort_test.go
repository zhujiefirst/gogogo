package sort

import (
	"fmt"
	// "math/rand"
	"testing"
)

func compareTo(lhs Comparable, rhs Comparable) int {
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

func genRandom(n int) []Comparable {
	var a []Comparable
	for i := 0; i < n; i++ {
		// e := rand.Int() % (2 * n)
		// e := i
		e := n - i
		a = append(a, e)
	}
	return a
}

func TestSort(*testing.T) {
	n := 10
	var c = make([]Comparable, n)
	var cc []Comparable
	var selectSort *SelectSort
	var insertionSort *InsertionSort

	cc = genRandom(n)

	copy(c[0:], cc[0:])
	selectSort = new(SelectSort)
	fmt.Println("before select sort...")
	selectSort.Show(c)
	selectSort.Sort(c, compareTo)
	fmt.Println("after select sort...")
	selectSort.Show(c)
	fmt.Printf("for select sort compare count=%v, exch count=%v\n", selectSort.compareCount, selectSort.exchCount)

	copy(c, cc)
	insertionSort = new(InsertionSort)
	fmt.Println("before insertion sort...")
	insertionSort.Show(c)
	insertionSort.Sort(c, compareTo)
	fmt.Println("after insertion sort...")
	insertionSort.Show(c)
	fmt.Printf("for insertion sort, compare count=%v, exch count=%v\n", insertionSort.compareCount, insertionSort.exchCount)
}