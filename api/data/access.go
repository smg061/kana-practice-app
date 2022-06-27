package data

import (
	"database/sql"
	"fmt"
	"os"
)

func GetDB() (*sql.DB, error) {
	var hostname string

	if (os.Getenv("ENV") == "production") {
		hostname = "pg"
	} else {
		hostname = "localhost"
	}

	var (
		host     = hostname
		port     = 5432 
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("pgx", psqlInfo)

	fmt.Println(psqlInfo)

	if err != nil {
		return nil, err
	}

	return db, nil

}
