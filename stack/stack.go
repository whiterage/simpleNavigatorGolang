package stack

import "errors"

var ErrEmptyStack = errors.New("stack is empty")

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

func (s *Stack) Push(v int) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (int, error) {
	top, err := s.Top()
	if err != nil {
		return top, err
	}
	s.items = s.items[:s.Size()-1]
	return top, nil
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}

	return s.items[s.Size()-1], nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
