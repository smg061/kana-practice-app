package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/smg061/kana-practice-app/api/controllers"
)

type app struct {
	db *sql.DB
}

func main() {
	err := loadSettings()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Logger.SetLevel(0)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.HEAD},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/kana", controllers.GetKana)
	e.GET("/test", controllers.GetStuff)
	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}

func loadSettings () error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}