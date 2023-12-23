package model

import "errors"

type Stack struct {
	values []string
}

func (s *Stack) Push(value string) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() (string, error) {
	if s.IsEmpty() {
		return "", errors.New("ERROR")
	}
	value := s.values[len(s.values) - 1]
	s.values = s.values[:len(s.values) - 1]
	return value, nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}
