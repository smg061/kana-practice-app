package controllers
import (
	"github.com/labstack/echo/v4"
	"github.com/smg061/kana-practice-app/api/models"
	"net/http"
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
