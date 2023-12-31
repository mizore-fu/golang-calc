package main

import (
	"app/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", hello)
	e.POST("/calc", calculate)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!!!")
}

//POST body: {"values": ["10", "+", "2", "*", "3"]}
func calculate(c echo.Context) error {
	formula := &model.Formula{}
	if err := c.Bind(formula); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	rpnFormula, err := formula.TranslateToRPNFormula()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	result, err := rpnFormula.Calculate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	resultStr := strconv.Itoa(result)
	values := strings.SplitAfter(resultStr, "")
	return c.JSON(http.StatusOK, &model.Response{Values: values})
}
