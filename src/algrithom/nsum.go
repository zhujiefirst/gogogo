package algrithom

import (
	// "log"
	"sort"
)

// 找出所有和为0的整数对的数量
func TwoSum(na []int32) (cnt int32) {
	l := len(na)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if na[i]+na[j] == 0 {
				cnt++
			}
		}
	}
	return
}

type nums []int32

func (a nums) Len() int           { return len(a) }
func (a nums) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a nums) Less(i, j int) bool { return a[i] < a[j] }

// 找出所有和为0的整数对的数量
// 先sort，再用bin-search方法寻找和为0的整数对
func TwoSumFast(na []int32) (cnt int32) {
	// log.Println("before sort:", na)
	sort.Sort(nums(na))
	// log.Println("before after:", na)

	for i, v := range na {
		idx := sort.Search(len(na), func(i int) bool { return na[i] >= -v })
		if idx != len(na) && i < idx && na[idx] == -v {
			// log.Printf("lv=%v[%v], rv=%v[%v]", v, i, na[idx], idx)
			cnt++
			for {
				idx++
				if idx < len(na) && na[idx] == -v {
					cnt++
					continue
				}
				break
			}
		}
	}

	return
}
