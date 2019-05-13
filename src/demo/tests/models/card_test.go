package models

import (
	"demo/models/card"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestCard(t *testing.T) {
	convey.Convey("testCard", t, func() {
		card := &card.Card{Index: 12, Shape: 3}
		convey.Convey("testCardGetValue", func() {
			convey.So(card.GetValue(), convey.ShouldEqual, 1203)
		})
	})
}

func TestCardSet(t *testing.T) {
	convey.Convey("CardSet", t, func() {
		cardSet := &card.CardSet{}
		convey.Convey("push", func() {
			cardSet.push
		})
		convey.Convey("sort", func() {

		})
	})
}
