package algrithom

// 找出所有和为0的整数对的数量
func TwoSum(na []int32) (cnt int32) {
	l := len(na)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if na[i]+na[j] == 0 {
				cnt++
			}
		}
	}
	return
}
