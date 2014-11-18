// 经典投色子问题。比如，投10次色子，10次色子的点数总共为50点，那么有多少种可能的投法(数学上可以用组合数学的多项式方法计算得到)
package algrithom

import (
// "log"
)

const (
	NonePoint = iota
	OnePoint
	TwoPoint
	ThreePoint
	FourPoint
	FivePoint
	SixPoint
	MinPoint = OnePoint
	MaxPoint = SixPoint
)

// 投times次数色子，得到points点数，有多少种可能投法
func ThrowDice(times int, points int) int {
	// log.Printf("throw dice %v times, getting total %v points", times, points)

	if times*MaxPoint < points {
		return 0
	}
	if times*MaxPoint == points {
		return 1
	}
	if times*MinPoint > points {
		return 0
	}
	if times*MinPoint == points {
		return 1
	}
	if times == 1 {
		return 1
	}
	nextTimes := times - 1
	return ThrowDice(nextTimes, points-OnePoint) + ThrowDice(nextTimes, points-TwoPoint) + ThrowDice(nextTimes, points-ThreePoint) + ThrowDice(nextTimes, points-FourPoint) + ThrowDice(nextTimes, points-FivePoint) + ThrowDice(nextTimes, points-SixPoint)
}
