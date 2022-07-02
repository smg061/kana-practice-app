package main

import (
	"log"
	"os"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/smg061/kana-practice-app/api/data"
	"github.com/smg061/kana-practice-app/api/models"
)

type app struct {
	KanaModel *models.KanaModel
}

func main() {

	err := loadSettings()
	if err != nil {
		log.Fatal(err)
	}

	// e := echo.New()


	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!\n")
	// })
	// e.GET("/kana", controllers.GetAllKana)
	
	//e.Logger.Fatal(e.Start(os.Getenv("PORT")))

	db, err := data.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	app := app {
		KanaModel: &models.KanaModel{
			DB: db,
		},
	}

	e := app.routes()
	e.Logger.SetLevel(0)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.HEAD},
	}))

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
	
}

func loadSettings () error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}