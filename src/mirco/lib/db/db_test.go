package db

import (
	"database/sql"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDb(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/game")
	if err != nil {
		t.Error("db error")
	}
	defer db.Close()

	Convey("testdb", t, func() {
		rows, _ := db.Query("select '123'")

		var user string
		for rows.Next() {
			err := rows.Scan(&user)
			if err != nil {
				t.Error(err)
			}
		}
	})

}
