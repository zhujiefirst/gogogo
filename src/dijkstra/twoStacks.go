package dijkstra

import (
	"container/list"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func oper(na *list.List, oa *list.List) {
	o := oa.Back()
	oa.Remove(o)
	a := na.Back()
	na.Remove(a)
	b := na.Back()
	na.Remove(b)
	switch o.Value.(string) {
	case "+":
		c := b.Value.(float32) + a.Value.(float32)
		na.PushBack(c)
		log.Println(b.Value, "+", a.Value, "=", c)
	case "-":
		c := b.Value.(float32) - a.Value.(float32)
		na.PushBack(c)
		log.Println(b.Value, "-", a.Value, "=", c)
	case "*":
		c := b.Value.(float32) * a.Value.(float32)
		na.PushBack(c)
		log.Println(b.Value, "*", a.Value, "=", c)
	case "/":
		c := b.Value.(float32) / a.Value.(float32)
		na.PushBack(c)
		log.Println(b.Value, "/", a.Value, "=", c)
	default:
		log.Fatalln("error oper=", o.Value.(string))
	}
}

func TwoStacks(s string) (v float32) {
	sa := strings.Split(s, " ")
	log.Println(sa)

	na := list.New()
	oa := list.New()

	for _, e := range sa {
		if m, _ := regexp.MatchString(`^\d[.\d]*$`, e); m {
			f, _ := strconv.ParseFloat(e, 32)
			na.PushBack(float32(f))
			log.Println("push into na", e)
		} else if m, _ := regexp.MatchString(`[\+\-\*/]`, e); m {
			oa.PushBack(e)
			log.Println("push into oa", e)
		} else if e == ")" {
			oper(na, oa)
		} else if e == "(" {
			log.Println("get '(', omit it")
		} else {
			log.Fatalln("error char=", e)
		}
	}

	if na.Len() != 1 || oa.Len() != 0 {
		log.Fatalln("error str=", s)
	}

	v = na.Back().Value.(float32)
	log.Println("result v=", v)

	return
}
