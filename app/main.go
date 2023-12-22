package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Formula struct {
	Formula []string `json:"formula"`
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

//POST body: {"formula": ["1", "2", "+"]}
func receiveInput(c echo.Context) error {
	jsonData := &Formula{}
	if err := c.Bind(jsonData); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	formulaValues := jsonData.Formula

	return c.JSON(http.StatusOK, formulaValues)
}
