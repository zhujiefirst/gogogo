package sort

import (
// "fmt"
)

type Comparable interface{}

type Compare func(lhs Comparable, rhs Comparable) (r int)
