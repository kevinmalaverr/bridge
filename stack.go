package bridge

type stack struct {
	array []interface{}
}

func (s *stack) push(data interface{}) {
	s.array = append(s.array, data)
}

func (s *stack) pop() interface{} {
	if len(s.array) > 0 {
		element := s.peek()
		s.array = s.array[:len(s.array)-1]
		return element
	}
	return nil
}

func (s *stack) peek() interface{} {
	return s.array[len(s.array)-1].(string)
}

func (s *stack) forEach(fn func(data interface{})) {
	for i := len(s.array) - 1; i >= 0; i-- {
		fn(s.array[i])
	}
}
