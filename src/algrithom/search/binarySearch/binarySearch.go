package binarySearch

import (
// "fmt"
)

type Key string
type Value string

type Node struct {
	key         Key
	value       Value
	left, right *Node
	N           int
}

type BST struct {
	root *Node
}

func (this *BST) Size() int {
	return this.size(this.root)
}

func (this *BST) size(n *Node) int {
	if n != nil {
		return n.N
	} else {
		return 0
	}
}

func (this *BST) Get(key Key) *Node {
	return this.get(this.root, key)
}

func (this *BST) get(n *Node, key Key) *Node {
	if n == nil {
		return nil
	}

	if n.key == key {
		return n
	}

	if n.key > key {
		return this.get(n.left, key)
	}

	if n.key < key {
		return this.get(n.right, key)
	}

	return nil
}

func (this *BST) Put(key Key, value Value) {
	this.root = this.put(this.root, key, value)
}

func (this *BST) put(n *Node, key Key, value Value) *Node {
	if n == nil {
		return &Node{key: key, value: value, N: 1}
	}

	if n.key == key {
		n.value = value
	} else if n.key < key {
		n.right = this.put(n.right, key, value)
	} else {
		n.left = this.put(n.left, key, value)
	}

	n.N = this.size(n.left) + this.size(n.right) + 1

	return n
}

func (this *BST) Min() Key {
	return this.min(this.root).key
}

func (this *BST) min(n *Node) *Node {
	if n.left != nil {
		return this.min(n.left)
	}
	return n
}

func (this *BST) Max() Key {
	return this.max(this.root).key
}

func (this *BST) max(n *Node) *Node {
	if n.right != nil {
		return this.max(n.right)
	}
	return n
}

func (this *BST) Floor(key Key) Key {
	if n := this.floor(this.root, key); n != nil {
		return n.key
	}

	return ""
}

func (this *BST) floor(n *Node, key Key) *Node {
	if n == nil {
		return nil
	}

	if key == n.key {
		return n
	}

	if key < n.key {
		return this.floor(n.left, key)
	}

	node := this.floor(n.right, key)
	if node == nil {
		return n
	} else {
		return node
	}

	return nil
}

func (this *BST) Ceiling(key Key) Key {
	if n := this.ceiling(this.root, key); n != nil {
		return n.key
	}

	return ""
}

func (this *BST) ceiling(n *Node, key Key) *Node {
	if n == nil {
		return nil
	}

	if key == n.key {
		return n
	}

	if key > n.key {
		return this.ceiling(n.right, key)
	}

	node := this.ceiling(n.left, key)
	if node == nil {
		return n
	} else {
		return node
	}

	return nil
}

func (this *BST) Select(k int) Key {
	if n := this.innerSelect(this.root, k); n != nil {
		return n.key
	}

	return ""
}

func (this *BST) innerSelect(n *Node, k int) *Node {
	if n == nil {
		return nil
	}

	if k <= 0 {
		return n
	}

	if n.left == nil {
		return this.innerSelect(n.right, k-1)
	}

	if n.left.N == k {
		return n
	}

	if n.left.N > k {
		return this.innerSelect(n.left, k)
	}

	if n.left.N < k {
		return this.innerSelect(n.right, k-n.left.N-1)
	}

	return nil
}

func (this *BST) Rank(key Key) int {
	return this.rank(this.root, key)
}

func (this *BST) rank(n *Node, key Key) int {
	if n == nil {
		return 0
	}

	if n.key == key {
		if n.left != nil {
			return n.left.N
		} else {
			return 0
		}
	}

	if n.key > key {
		return this.rank(n.left, key)
	}

	if n.key < key {
		if n.left != nil {
			return n.left.N + 1 + this.rank(n.right, key)
		} else {
			return 1 + this.rank(n.right, key)
		}
	}

	return 0
}

func New() *BST {
	return &BST{}
}
