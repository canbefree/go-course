package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	TYPE_MYSQL = "mysql"
)

type DB struct {
	db sql.DB
}

func (database *DB) Use(dblink sql.DB) {
	database.db = dblink
}

/**
usage:
	DB.insert();
	import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

	db.insert();
	db.query()
	db.one();

*/
