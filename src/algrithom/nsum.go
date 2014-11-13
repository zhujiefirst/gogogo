package algrithom

import (
	"log"
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

// 找出所有和为0的三元整数对的数量
func ThreeSum(na []int32) (cnt int32) {
	l := len(na)
	// log.Println(na)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				// log.Printf("%v+%v+%v=%v", na[i], na[j], na[k], na[i]+na[j]+na[k])
				if na[i]+na[j]+na[k] == 0 {
					cnt++
				}
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
	log.Println("before sort:", na)
	sort.Sort(nums(na))
	log.Println("after sort:", na)

	for i, v := range na {
		idx := sort.Search(len(na), func(i int) bool { return na[i] >= -v })
		log.Printf("na[%v]=%v, idx=%v", i, v, idx)
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

func ThreeSumFast(na []int32) (cnt int32) {
	// log.Println("before sort:", na)
	sort.Sort(nums(na))
	// log.Println("after sort:", na)

	for i, vi := range na {
		for j, vj := range na[i:] {
			v := vi + vj
			idx := sort.Search(len(na), func(i int) bool { return na[i] >= -v })
			if idx != len(na) && func(a, b int) int {
				if a > b {
					return a
				}
				return b
			}(i, j) < idx && na[idx] == -v {
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
	}

	return
}
