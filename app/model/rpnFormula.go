package model

import (
	"app/functions"
	"errors"
	"strconv"
)

func calcTwoValues(firstNumber int, secondNumber int, operator string) (int, error) {
	var result int
	switch operator {
		case "+":
			result = firstNumber + secondNumber
		case "-":
			result = firstNumber - secondNumber
		case "*":
			result = firstNumber * secondNumber
		case "/":
			if secondNumber == 0 {
				return 0, errors.New("division by 0")
			}
			result = firstNumber / secondNumber
	}
	return result, nil
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
			result, err := calcTwoValues(firstNumber, secondNumber, operator)
			if err != nil {
				return 0, err
			}
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
