package stack

type Stack struct {
	data []any
}

type Iterable interface {
	Next() (interface{}, bool)
}

func New(value any) Stack {
	var s Stack

	switch val := value.(type) {
	case []any:
		s.data = append(s.data, val...)
	case string:
		for _, v := range val {
			s.data = append(s.data, string(v))
		}
	}

	return s
}

func (s *Stack) Push(val any) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() any {
	if s.IsEmpty() {
		return nil
	}

	deleted := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return deleted
}

func (s *Stack) Peek() any {
	if s.IsEmpty() {
		return nil
	}
	return s.data[len(s.data)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
