package algrithom

import (
// "fmt"
)

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

func (u *UF) Unit(p int, q int) {
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

func (u *UF) Find(p int) int {
	return u.id[p]
}

func (u *UF) Connected(p int, q int) bool {
	return u.Find(p) == u.Find(q)
}

func (u *UF) Count() int {
	return u.count
}
