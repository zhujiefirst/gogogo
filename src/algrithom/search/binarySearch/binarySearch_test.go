package binarySearch

import (
	// "fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	bst := New()

	// empty btree
	if bst.Size() != 0 {
		t.Errorf("[EXCEPT]=0\t[FACTOR]=%v", bst.Size())
	}

	if bst.Get("return_nil") != nil {
		t.Errorf("[EXCEPT]=nil\t[FACTOR]=%v", bst.Get("return_nil"))
	}

	// put 1 elem
	bst.Put("1", "first_value")
	if bst.Size() != 1 {
		t.Errorf("[EXCEPT]=1\t[FACTOR]=%v", bst.Size())
	}

	// put 1 elem
	bst.Put("2", "second_value")
	if bst.Size() != 2 {
		t.Errorf("[EXCEPT]=2\t[FACTOR]=%v", bst.Size())
	}

	// put 1 elem
	bst.Put("3", "third_value")
	if bst.Size() != 3 {
		t.Errorf("[EXCEPT]=3\t[FACTOR]=%v", bst.Size())
	}

	// put 1 elem
	bst.Put("3", "third_value_2")
	if bst.Size() != 3 {
		t.Errorf("[EXCEPT]=3\t[FACTOR]=%v", bst.Size())
	}

	// get elem which key=2
	node := bst.Get("2")
	if node == nil || node.value != "second_value" {
		t.Errorf("[EXCEPT]=second_value\t[FACTOR]=%v", bst.Get("2"))
	}

	// get elem which is not exist
	node = bst.Get("nil_node")
	if node != nil {
		t.Errorf("[EXCEPT]=nil\t[FACTOR]=not nil")
	}

	if bst.Max() != "3" {
		t.Errorf("[EXCEPT]=3\t[FACTOR]=%v", bst.Max())
	}

	if bst.Min() != "1" {
		t.Errorf("[EXCEPT]=1\t[FACTOR]=%v", bst.Min())
	}

	// put 1 elem
	bst.Put("5", "five_value")
	if bst.Size() != 4 {
		t.Errorf("[EXCEPT]=4\t[FACTOR]=%v", bst.Size())
	}

	if bst.Floor("4") != "3" {
		t.Errorf("[EXCEPT]=3\t[FACTOR]=%v", bst.Floor("4"))
	}

	if bst.Floor("3") != "3" {
		t.Errorf("[EXCEPT]=3\t[FACTOR]=%v", bst.Floor("3"))
	}

	if bst.Floor("0") != "" {
		t.Errorf("[EXCEPT]=\"\"\t[FACTOR]=%v", bst.Floor("0"))
	}

	if bst.Ceiling("6") != "" {
		t.Errorf("[EXCEPT]=\"\"\t[FACTOR]=%v", bst.Ceiling("6"))
	}

	if bst.Ceiling("4") != "5" {
		t.Errorf("[EXCEPT]=\"5\"\t[FACTOR]=%v", bst.Ceiling("4"))
	}

	if bst.Select(1) != "2" {
		t.Errorf("[EXCEPT]=\"2\"\t[FACTOR]=%v", bst.Select(1))
	}

	if bst.Rank("1") != 0 {
		t.Errorf("[EXCEPT]=0\t[FACTOR]=%v", bst.Rank("1"))
	}

	if bst.Rank("3") != 2 {
		t.Errorf("[EXCEPT]=2\t[FACTOR]=%v", bst.Rank("3"))
	}
}
