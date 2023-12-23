package main

import (
	"app/model"
	"errors"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CalcTwoValues(firstNumber int, secondNumber int, operator string) int {
	switch operator {
		case "+":
			return firstNumber + secondNumber
		case "-":
			return firstNumber - secondNumber
		case "*":
			return firstNumber * secondNumber
		case "/":
			return firstNumber / secondNumber
		default:
			return 0
	}
}

type Formula struct {
	Values []string `json:"values"`
}

func (f *Formula) Calculate() (int, error) {
	stack := &model.Stack{}
	values := f.Values
	regexpNumber := regexp.MustCompile(`[0-9]`)
	regexpOperator := regexp.MustCompile(`[\+\-\*\/]`)
	for i := 0; i < len(values); i++ {
		value := values[i]
		if regexpNumber.MatchString(value) {
			stack.Push(value)
			continue
		}
		if regexpOperator.MatchString(value) {
			operator := value
			secondStr, err := stack.Pop()
			if err != nil {
				return 0, err
			}
			firstStr, err := stack.Pop()
			if err != nil {
				return 0, err
			}
			secondNumber, err := strconv.Atoi(secondStr)
			if err != nil {
				return 0, err
			}
			firstNumber, err := strconv.Atoi(firstStr)
			if err != nil {
				return 0, err
			}
			result := CalcTwoValues(firstNumber, secondNumber, operator)
			stack.Push(strconv.Itoa(result))
			continue
		}
		return 0, errors.New("ERROR")
	}
	resultStr, err := stack.Pop()
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(resultStr)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.POST("/calc", receiveInput)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!!!")
}

//POST body: {"values": ["1", "2", "+"]}
func receiveInput(c echo.Context) error {
	formula := &Formula{}
	if err := c.Bind(formula); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	value, err := formula.Calculate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, &model.Response{Value: value})
}
