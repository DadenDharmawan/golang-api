package apps

import (
	"database/sql"
	"time"

	"github.com/DadenDharmawan/api-go/helper"
)

func ConnectionToDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/api_go")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}