package controllers


import (

	"github.com/labstack/echo/v4"
	"net/http"

)

type Kana struct {
	Representation string `json:"representation"`
	Romaji string `json:"romaji"`

}

func GetKana(c echo.Context) (err error) {
	
	kana := []Kana{
		{"ひ", "hi"},
		{"は", "hi"},
	}

	return c.JSON(http.StatusOK, kana)
	
}