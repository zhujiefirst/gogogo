package algrithom

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

const (
	N = 5
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestTwoSum(*testing.T) {
	var na []int32
	for i := 0; i < N; i++ {
		na = append(na, rand.Int31n(N)-N/2)
	}

	tb := time.Now()
	log.Printf("在数量为%v的随机整数数组中，和为0的整数对数量为%v, expend=%v", len(na), TwoSum(na), time.Now().Sub(tb))

	tb = time.Now()
	log.Printf("在数量为%v的随机整数数组中，和为0的整数对数量为%v, expend=%v", len(na), TwoSumFast(na), time.Now().Sub(tb))
}

func TestThreeSum(*testing.T) {
	var na []int32
	for i := 0; i < N; i++ {
		na = append(na, rand.Int31n(N)-N/2)
	}

	tb := time.Now()
	log.Printf("在数量为%v的随机整数数组中，和为0的三元整数对数量为%v, expend=%v", len(na), ThreeSum(na), time.Now().Sub(tb))

	tb = time.Now()
	log.Printf("在数量为%v的随机整数数组中，和为0的三元整数对数量为%v, expend=%v", len(na), ThreeSumFast(na), time.Now().Sub(tb))
}
