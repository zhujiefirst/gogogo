package structData

type None struct{}

type Set struct {
	set map[interface{}]None
}

func (s *Set) InSet(e interface{}) bool {
	_, ok := s.set[e]
	return ok
}

func (s *Set) Insert(e interface{}) bool {
	in := s.InSet(e)
	s.set[e] = None{}
	return in
}

func (s *Set) Clear() {
	s.set = make(map[interface{}]None)
}

func (s *Set) Size() int {
	return len(s.set)
}

func (s *Set) Range() map[interface{}]None {
	return s.set
}
