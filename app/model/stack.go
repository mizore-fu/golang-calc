package model

import "errors"

type Stack struct {
	values []string
}

func (s *Stack) Push(value string) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() (string, error) {
	if err := s.IsNoData(); err != nil {
		return "", err
	}
	value := s.values[len(s.values) - 1]
	s.values = s.values[:len(s.values) - 1]
	return value, nil
}

func (s *Stack) IsNoData() error {
	if len(s.values) == 0 {
		return errors.New("ERROR")
	}
	return nil
}
