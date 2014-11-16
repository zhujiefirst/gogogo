// 生成无重复随机数
package algrithom

import (
	"container/list"
	// "fmt"
	"log"
	"math/rand"
	"time"
)

// 暴力方法
func forceMethod(N int) {
	tb := time.Now()
	na := make([]int, 0, N)
Loop:
	for {
		n := rand.Intn(2 * N)
		for _, v := range na {
			if n == v {
				continue Loop
			}
		}
		na = append(na, n)
		if len(na) == N {
			break
		}
	}
	// fmt.Println(na)
	log.Printf("forceMethod:%v", time.Now().Sub(tb))
}

// 预先生成范围内所有数存在array中，随机生成位置数，每获得一个数后将其删除
func arrayMethod(N int) {
	tb := time.Now()
	na := make([]int, 0, N)
	pna := make([]int, 0, 2*N)
	for i := 0; i < 2*N; i++ {
		pna = append(pna, i)
	}
	for i := 0; i < N; i++ {
		p := rand.Intn(2*N - i)
		na = append(na, pna[p])
		pna = append(pna[:p], pna[p+1:]...)
	}
	// fmt.Println(na)
	log.Printf("arrayMethod:%v", time.Now().Sub(tb))
}

// 预先生成范围内所有数存在list中，随机生成位置数，每获得一个数后将其删除
func listMethod(N int) {
	tb := time.Now()
	na := make([]int, 0, N)
	pna := list.New()
	for i := 0; i < 2*N; i++ {
		pna.PushBack(i)
	}
	for i := 0; i < N; i++ {
		p := rand.Intn(2*N - i)
		na = append(na, func(l *list.List, pos int) int {
			e := l.Front()
			for j := 0; j < pos; j++ {
				e = e.Next()
			}
			l.Remove(e)
			return e.Value.(int)
		}(pna, p))
	}
	// fmt.Println(na)
	log.Printf("listMethod:%v", time.Now().Sub(tb))
}

// 预先生成范围内所有数存在array中，随机生成位置数，每获得一个数后将与最后一位数交换
func betterMethod(N int) {
	tb := time.Now()
	na := make([]int, 0, N)
	pna := make([]int, 0, 2*N)
	for i := 0; i < 2*N; i++ {
		pna = append(pna, i)
	}
	for i := 0; i < N; i++ {
		p := rand.Intn(2*N - i)
		na = append(na, pna[p])
		pna[p], pna[2*N-i-1] = pna[2*N-i-1], pna[p]
	}
	log.Printf("betterMethod:%v", time.Now().Sub(tb))
}
