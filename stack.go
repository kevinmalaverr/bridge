package bridge

type Stack struct {
	array []interface{}
}

func (s *Stack) push(data interface{}) {
	s.array = append(s.array, data)
}

func (s *Stack) pop() interface{} {
	if len(s.array) > 0 {
		element := s.peek()
		s.array = s.array[:len(s.array)-1]
		return element
	}
	return nil
}

func (s *Stack) peek() interface{} {
	return s.array[len(s.array)-1]
}

func (s *Stack) forEach(fn func(data interface{})) {
	for i := len(s.array) - 1; i >= 0; i-- {
		fn(s.array[i])
	}
}
