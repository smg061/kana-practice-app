package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/smg061/kana-practice-app/api/data"
	"github.com/smg061/kana-practice-app/api/models"
)
const  (
	Katakana = "katakana"
	Hiragana = "hiragana"
	Kanji = "kanji"
)

func GetKana(c echo.Context) (err error) {
	kana := []models.Kana{
		{Representation: "ひ", Romaji: "hi", Classification: Hiragana},
		{Representation: "は", Romaji: "ha", Classification: Hiragana},
	}
	return c.JSON(http.StatusOK, kana)
}

type Test struct {
	Id int `json:"id"`
	String string `json:"string"`
}
func GetStuff(c echo.Context) (err error) {

	db, err := data.GetDB()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "`message: something went wrong with the database")
	}
	rows, err := db.Query(`SELECT * FROM test`)
	if err != nil {
		fmt.Println(err)
		//c.Echo().Logger.Error("error accessing data")
		c.Logger().Error("error accessing data")
		return c.JSON(http.StatusNotFound, "`message: something went wrong with retrieving the rows`")
	}
	defer rows.Close()
	testStructs := []Test{}
	for rows.Next() {
		var c Test
		err := rows.Scan(&c.Id, &c.String)
		if err != nil {
			continue
		}
		testStructs = append(testStructs, c)
	}
	c.Logger().Info("data access succcessfull")
	return c.JSON(http.StatusOK, testStructs)
}