package algrithom

import (
// "fmt"
)

type FindFunc func(int) int

type UFI interface {
	Init(n int)
	Unit(p int, q int)
	Find(p int) int
	Connected(find FindFunc, p int, q int) bool
	Count() int
}

type UF struct {
	id    []int
	count int
}

func (u *UF) Init(n int) {
	u.count = n
	u.id = make([]int, 0, n)
	for i := 0; i < n; i++ {
		u.id = append(u.id, i)
	}
}

func (u *UF) Count() int {
	return u.count
}

func (u *UF) Connected(find FindFunc, p int, q int) bool {
	return find(p) == find(q)
}

type UFQuickFind struct {
	UF
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

type UFQuickUnit struct {
	UF
}

func (u *UFQuickUnit) Unit(p int, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)

	if pRoot == qRoot {
		return
	}

	u.id[pRoot] = qRoot
	u.count--
}

func (u *UFQuickUnit) Find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

type UFWeightQuickUnit struct {
	UF
	sz []int
}

func (u *UFWeightQuickUnit) Init(n int) {
	u.count = n
	u.id = make([]int, 0, n)
	u.sz = make([]int, 0, n)
	for i := 0; i < n; i++ {
		u.id = append(u.id, i)
		u.sz = append(u.sz, 1)
	}
}

func (u *UFWeightQuickUnit) Unit(p int, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)

	if pRoot == qRoot {
		return
	}

	if u.sz[pRoot] < u.sz[qRoot] {
		u.id[pRoot] = qRoot
		u.sz[qRoot] = u.sz[qRoot] + u.sz[pRoot]
	} else {
		u.id[qRoot] = pRoot
		u.sz[pRoot] = u.sz[pRoot] + u.sz[qRoot]
	}
	u.count--
}

func (u *UFWeightQuickUnit) Find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}
