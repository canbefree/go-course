package helper

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinQuantity(t *testing.T) {
	Convey("min", t, func() {
		subject := minQuantity(31282)
		So(subject, ShouldEqual, 32768)
	})
}
