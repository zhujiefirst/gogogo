package priority_queue

import (
// "fmt"
)

type Comparable interface{}

type Compare func(lhs Comparable, rhs Comparable) int
