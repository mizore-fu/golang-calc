package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Formula struct {
	Values []string `json:"values"`
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
	formulaValues := formula.Values

	return c.JSON(http.StatusOK, formulaValues)
}
