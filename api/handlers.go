package main
import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/smg061/kana-practice-app/api/data"
)
const  (
	Hiragana = iota
	Katakana 
	Kanji
)

func (app *app)GetAllKana(c echo.Context) (err error) {

	db,err := data.GetDB()

	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	statement := "SELECT representation, romaji, classification, initial FROM kana"

	rows, err := db.Query(statement)

	if err != nil {
		c.Logger().Error(err)
	}
	defer rows.Close()

 	kanaResult,err := app.KanaModel.GetAllKana()
	if err != nil {
		c.Logger().Error(err)
	}
	return c.JSON(http.StatusOK, kanaResult)
}


func InsertKanaRecords (c echo.Context) (err error) {

	insertRows, err := data.ReadKanaRecords()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	db, err := data.GetDB()
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