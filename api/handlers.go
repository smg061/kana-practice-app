package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/smg061/kana-practice-app/api/data"
)
const  (
	Hiragana = iota
	Katakana 
	Kanji
)

type KanaRequest map[string]interface{}

func (app *app) GetAllKana(c echo.Context) (err error) {
	
 	kanaResult,err := app.KanaModel.GetAllKana()
	if err != nil {
		c.Logger().Error(err)
	}
	return c.JSON(http.StatusOK, kanaResult)
}

func (app *app) GetKanaByClass(c echo.Context) (err error) {
	req := KanaRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	// json numbers that come from req are float64 values 
	filter := req["filter"].(float64)

	if err != nil {
		// c.Logger().Error("Error")
		return c.JSON(http.StatusInternalServerError, err)
	}
	kanaResult, err := app.KanaModel.GetByClass(filter)
	if err != nil {
		c.Logger().Error(err)
	}
	return c.JSON(http.StatusOK, kanaResult)
}

func (app *app) GetKanaByInitial (c echo.Context) (err error) {
	req:= KanaRequest{}
	if err:= c.Bind(&req); err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusBadRequest, err)
	}

	filter, ok := req["filter"].(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, "filter must be a string or something castable as a string")
	}

	kana, err := app.KanaModel.GetByInitial(filter)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong when getting results")
	}

	return c.JSON(http.StatusOK, kana)
}

func (app *app) GetKanaByInitialQuery(c echo.Context) (err error) {
	initial := c.Param("initial")
	kana, err := app.KanaModel.GetByInitial(initial)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, "Something went wrong when getting results")
	}
	if len(kana) == 0 {
		return c.JSON(http.StatusNotFound, fmt.Sprintf("Could not find data with '%s' as the initial", initial))
	}
	return c.JSON(http.StatusOK, kana)
}
func (app *app) GetKanaByClassQuery(c echo.Context) (err error) {
	class := c.Param("type")

	if class == "" {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Please provide a valid parameter"))
	}
	var kanaType float64
	
	if class == "hiragana" {
		kanaType = 0
	} else {
		kanaType = 1.0
	}
	kana, err := app.KanaModel.GetByClass(kanaType)
	if err != nil {
		return c.JSON(http.StatusAccepted, fmt.Sprintf("Something went wrong: %v", err))
	}
	return c.JSON(http.StatusOK, kana)
}

func InsertKanaRecords (c echo.Context) (err error) {

	insertRows, err := data.ReadKanaRecords()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	db, err := data.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	for _, row := range insertRows {

		statement := "INSERT INTO kana(representation, romaji, classification, initial) VALUES ($1, $2, $3, $4)"

		res, err := db.Exec(statement, &row.Representation, &row.Romaji, &row.Classification, &row.Initial)

		c.Logger().Info(res)
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, err)
		}
	}
	return c.JSON(http.StatusAccepted, fmt.Sprintf("inserted %d rows into kana table", len(insertRows)))
}