package search

import (
	"container/list"
	// "log"
)

type Key string
type Value int

type node struct {
	key   Key
	value Value
}

type Sequence struct {
	list.List
}

func (this *Sequence) get(key Key) (value Value, err bool) {
	err = false

	for e := this.Front(); e != nil; e = e.Next() {
		if n, ok := e.Value.(node); ok {
			if key == n.key {
				value = n.value
				err = true
				return
			}
		}
	}

	return
}

func (this *Sequence) put(key Key, value Value) {
	for e := this.Front(); e != nil; e = e.Next() {
		if n, ok := e.Value.(node); ok {
			if key == n.key {
				this.Remove(e)
				break
			}
		}
	}

	e := node{key, value}
	this.PushFront(e)
}
