package scenes

import (
	"errors"
	"sync"
)

type Stack struct {
	lock sync.Mutex
	s    []interface{}
}

func (s Stack) Peek() interface{} {
	if s.Len() == 0 {
		return nil
	}

	return s.s[s.Len()-1]
}

func (s Stack) Len() int {
	return len(s.s)
}

func (s Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return nil, errors.New("empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func newStack() *Stack {
	return &Stack{sync.Mutex{}, make([]interface{}, 0)}
}
