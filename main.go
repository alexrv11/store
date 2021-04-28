package main

import (
	"fmt"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!2")
	})

	e.GET("/div", DivideHandler)
	e.Logger.Fatal(e.Start(":1323"))
}

func DivideHandler(c echo.Context) error {

	a := 1
	b := 0

	res := fmt.Sprintf("%d", a/b)

	return c.String(http.StatusOK, res)
}