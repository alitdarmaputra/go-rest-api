package app

import (
	"database/sql"
	"time"

	"github.com/alitdarmaputra/belajar-golang-rest-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root123@tcp(localhost:3306)/belajar_golang_db_migrate")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Second)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
