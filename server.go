package main

import (
	"net/http"
	"github.com/labstack/echo"
	_ "gopkg.in/cq.v1"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "HELLO FROM API")
	})

	e.POST("/users", createUser)
	e.Logger.Fatal(e.Start(":8000"))
}