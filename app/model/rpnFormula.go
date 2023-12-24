package model

import (
	"app/functions"
	"errors"
	"strconv"
)

func calcTwoValues(firstNumber int, secondNumber int, operator string) int {
	var result int
	switch operator {
		case "+":
			result = firstNumber + secondNumber
		case "-":
			result = firstNumber - secondNumber
		case "*":
			result = firstNumber * secondNumber
		case "/":
			result = firstNumber / secondNumber
	}
	return result
}

type RPNFormula struct {
	values []string
}

func (rf *RPNFormula) Calculate() (int, error) {
	stack := &Stack{}
	values := rf.values
	for i := 0; i < len(values); i++ {
		value := values[i]
		if functions.IsInteger(value) {
			stack.Push(value)
			continue
		}
		if functions.IsOperator(value) {
			operator := value
			secondStr, err := stack.Pop()
			if err != nil {
				return 0, err
			}
			firstStr, err := stack.Pop()
			if err != nil {
				return 0, err
			}
			secondNumber, _ := strconv.Atoi(secondStr)
			firstNumber, _ := strconv.Atoi(firstStr)
			result := calcTwoValues(firstNumber, secondNumber, operator)
			stack.Push(strconv.Itoa(result))
			continue
		}
		return 0, errors.New("ERROR")
	}
	resultStr, err := stack.Pop()
	if err != nil {
		return 0, err
	}
	result, _ := strconv.Atoi(resultStr)
	return result, nil
}
