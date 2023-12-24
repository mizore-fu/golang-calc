package main

import (
	"app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.POST("/calc", receiveInput)
	e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world!!!")
}

//POST body: {"values": ["1", "+", "2", "*", "3"]}
func receiveInput(c echo.Context) error {
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

	return c.JSON(http.StatusOK, &model.Response{Value: result})
}
