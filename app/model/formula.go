package model

import (
	"app/functions"
	"errors"
)

func isHighPriority(value string) bool {
	return value == "*" || value == "/"
}

func isLowerPriority(firstValue string, secondValue string) bool {
	if isHighPriority(firstValue) {
		return false
	}
	if isHighPriority(secondValue) {
		return true
	}
	return false
}

type Formula struct {
	Values []string `json:"values"`
}

func (f *Formula) TranslateToRPNFormula() (*RPNFormula, error) {
	stack := &Stack{}
	values := f.Values
	rpnValues := []string {}
	for i := 0; i < len(values); i++ {
		value := values[i]
		if functions.IsInteger(value) {
			rpnValues = append(rpnValues, value)
			continue
		}

		if !functions.IsOperator(value)  {
			return nil, errors.New("ERROR")
		}
		if stack.IsEmpty() {
			stack.Push(value)
			continue
		}
		stackTop, _ := stack.Pop()
		if isLowerPriority(value, stackTop) {
			rpnValues = append(rpnValues, stackTop)
			stack.Push(value)
			continue
		}
		stack.Push(stackTop)
		stack.Push(value)
	}
	for !stack.IsEmpty() {
		stackTop, _ := stack.Pop()
		rpnValues = append(rpnValues, stackTop)
	}
	rpnFormula := &RPNFormula{values: rpnValues}
	return rpnFormula, nil
}
