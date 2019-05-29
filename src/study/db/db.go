package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/game")
	return db, err
}
