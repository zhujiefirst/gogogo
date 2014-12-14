package sort

import (
	"fmt"
	"math/rand"
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
		e := rand.Int() % (2 * n)
		// e := i
		// e := n - i
		a = append(a, e)
	}
	return a
}

func TestSort(*testing.T) {
	n := 10
	var c = make([]Comparable, n)
	var cc []Comparable

	cc = genRandom(n)

	selectSortAl := func() {
		copy(c[0:], cc[0:])
		selectSort := new(SelectSort)
		fmt.Println("before select sort...")
		selectSort.Show(c)
		selectSort.Sort(c, compareTo)
		fmt.Println("after select sort...")
		selectSort.Show(c)
		fmt.Printf("for select sort compare count=%v, exch count=%v\n", selectSort.compareCount, selectSort.exchCount)
	}

	insertionSortAl := func() {
		copy(c, cc)
		insertionSort := new(InsertionSort)
		fmt.Println("before insertion sort...")
		insertionSort.Show(c)
		insertionSort.Sort(c, compareTo)
		fmt.Println("after insertion sort...")
		insertionSort.Show(c)
		fmt.Printf("for insertion sort, compare count=%v, exch count=%v\n", insertionSort.compareCount, insertionSort.exchCount)
	}

	shellSortAl := func() {
		copy(c, cc)
		shellSort := new(ShellSort)
		fmt.Println("before shell sort...")
		shellSort.Show(c)
		shellSort.Sort(c, compareTo)
		fmt.Println("after shell sort...")
		shellSort.Show(c)
		fmt.Printf("for shell sort, compare count=%v, exch count=%v\n", shellSort.compareCount, shellSort.exchCount)
	}

	bubbleSortAl := func() {
		copy(c, cc)
		bubbleSort := new(BubbleSort)
		fmt.Println("before bubble sort...")
		bubbleSort.Show(c)
		bubbleSort.Sort(c, compareTo)
		fmt.Println("after bubble sort...")
		bubbleSort.Show(c)
		fmt.Printf("for bubble sort, compare count=%v, exch count=%v\n", bubbleSort.compareCount, bubbleSort.exchCount)
	}

	mergeSortAl := func() {
		copy(c, cc)
		mergeSort := new(MergeSort)
		fmt.Println("before merge sort...")
		mergeSort.Show(c)
		mergeSort.Sort(c, compareTo)
		fmt.Println("after merge sort...")
		mergeSort.Show(c)
		fmt.Printf("for merge sort, compare count=%v, exch count=%v\n", mergeSort.compareCount, mergeSort.exchCount)
	}

	if true {
		selectSortAl()
		insertionSortAl()
		shellSortAl()
		bubbleSortAl()
		mergeSortAl()
	}
}
