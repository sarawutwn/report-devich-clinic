package database

import (
	"backend-app/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func Connect() error {
	fmt.Println(config.GetEnvConfig("DATABASE_URL"))
	db, err := sql.Open("postgres", config.GetEnvConfig("DATABASE_URL"))
	if err != nil {
		fmt.Println("connect fail" + err.Error())
		return err
	}

	DB = db
	fmt.Println("Connected to database")
	return nil

}
