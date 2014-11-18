package algrithom

import (
	"log"
	"testing"
)

func TestThrowDice(*testing.T) {
	times := 10
	points := 50
	log.Printf("Throw dice %v times, getting %v points, total %v counts", times, points, ThrowDice(times, points))
}
