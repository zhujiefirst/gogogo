package algrithom

import (
// "fmt"
)

type UFI interface {
	Init(n int)
	Unit(p int, q int)
	Find(p int) int
	Connected(p int, q int) bool
	Count() int
}

type UF struct {
	id    []int
	count int
}

type UFQuickFind struct {
	UF
}

func (u *UFQuickFind) Init(n int) {
	u.count = n
	u.id = make([]int, 0, n)
	for i := 0; i < n; i++ {
		u.id = append(u.id, i)
	}
}

func (u *UFQuickFind) Unit(p int, q int) {
	pID := u.Find(p)
	qID := u.Find(q)

	if pID == qID {
		return
	}

	for i := 0; i < len(u.id); i++ {
		if u.id[i] == pID {
			u.id[i] = qID
		}
	}

	u.count--
}

func (u *UFQuickFind) Find(p int) int {
	return u.id[p]
}

func (u *UFQuickFind) Connected(p int, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UFQuickFind) Count() int {
	return u.count
}
