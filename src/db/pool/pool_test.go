package pool

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPool(t *testing.T) {
	Convey("test pool", t, func() {

		db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/game")
		So(err, ShouldNotBeNil)
		db.Query("select * from user")
	})
}


