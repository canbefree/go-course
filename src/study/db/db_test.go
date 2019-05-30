package db

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSuite(t *testing.T) {
	Convey("test db", t, func() {
		db, err := NewConnect()
		Convey("get conn", func() {
			So(err, ShouldBeNil)
		})

		Convey("create table", func() {
			createSql := "create table if not exists test (id int(10) primary key, name varchar(20) not null);"
			stmts, err := db.Prepare(createSql)
			So(err, ShouldBeNil)
			result, err := stmts.Exec()

			So(err, ShouldBeNil)

			affected, _ := result.RowsAffected()

			So(affected, ShouldEqual, 0)
		})

		Convey("drop table", func() {
			dropSql := "drop table if exists test"
			result, err := db.Exec(dropSql)
			So(err, ShouldBeNil)
			affected, _ := result.RowsAffected()
			So(affected, ShouldEqual, 0)
		})

		Convey("test convey", func(c C) {
			c.So(4, ShouldBeBetween, 3, 4)
		})
	})
}
