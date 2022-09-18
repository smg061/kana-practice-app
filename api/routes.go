package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func (app *app) routes() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	e.GET("/kana", app.GetAllKana)
	e.POST("/kanaByClass", app.GetKanaByClass)
	e.POST("/kanaByInitial", app.GetKanaByInitial)
	e.GET("/kana/:initial", app.GetKanaByInitialQuery)
	e.GET("/kana/byclass/:class", app.GetKanaByClassQuery)
	return e

}