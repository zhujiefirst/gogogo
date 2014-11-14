package algrithom

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

const (
	N = 2000
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

type NoNone struct {
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func genRandom(n int32) []int32 {
	var na []int32
	flag := make(map[int32]struct{})
	for {
		rn := rand.Int31n(2 * N)
		if _, ok := flag[rn]; !ok {
			na = append(na, rn-2*N/2)
			flag[rn] = NoNone{}
			if int32(len(na)) == n {
				break
			}
		}
	}
	return na
}

func TestTwoSum(*testing.T) {
	na := genRandom(N)
	log.Printf("gen %v ints", len(na))

	tb := time.Now()
	log.Printf("在数量为%v的随机整数数组中，和为0的整数对数量为%v, expend=%v", len(na), TwoSum(na), time.Now().Sub(tb))

	tb = time.Now()
	log.Printf("在数量为%v的随机整数数组中，和为0的整数对数量为%v, expend=%v", len(na), TwoSumFast(na), time.Now().Sub(tb))
}

func TestThreeSum(*testing.T) {
	na := genRandom(N)
	log.Printf("gen %v ints", len(na))

	tb := time.Now()
	log.Printf("在数量为%v的随机整数数组中，和为0的三元整数对数量为%v, expend=%v", len(na), ThreeSum(na), time.Now().Sub(tb))

	tb = time.Now()
	log.Printf("在数量为%v的随机整数数组中，和为0的三元整数对数量为%v, expend=%v", len(na), ThreeSumFast(na), time.Now().Sub(tb))
}
