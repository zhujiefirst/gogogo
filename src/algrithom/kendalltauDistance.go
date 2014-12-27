package algrithom

import (
// "fmt"
)

// 采用插入排序，记录元素交换次数
// 交换次数即不按顺序的元素对数
func insertSort(a []int) int {
	var exchCnt = 0
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				var tmp = a[j]
				a[j] = a[j-1]
				a[j-1] = tmp
				exchCnt++
			}
		}
	}

	return exchCnt
}

// KendalltauDistance 计算两对序列的Kendall tau距离
func KendalltauDistance(a []int, b []int) int {
	var c = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				c[j] = i
			}
		}
	}

	return insertSort(c)
}

// func main() {
// 	a := []int{0, 3, 1, 6, 2, 5, 4}
// 	b := []int{1, 0, 3, 6, 4, 2, 5}
// 	fmt.Println(KendalltauDistance(a, b))
// }
