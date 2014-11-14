package structData

import (
	"log"
	"testing"
)

func TestSet(*testing.T) {
	set := Set{set: make(map[interface{}]None)}
	v := "elem"

	if set.InSet(v) == true {
		log.Panic("")
	}
	if set.Insert(v) == true {
		log.Panic("")
	}
	if set.Size() != 1 {
		log.Panic("")
	}
	if set.InSet(v) == false {
		log.Panic("")
	}
	if set.Insert(v) == false {
		log.Panic("")
	}
	if set.Size() != 1 {
		log.Panic("")
	}
	for e, _ := range set.Range() {
		if e != v {
			log.Panic("")
		}
	}
	set.Clear()
	if set.InSet(v) == true {
		log.Panic("")
	}
	if set.Size() != 0 {
		log.Panic("")
	}
}
