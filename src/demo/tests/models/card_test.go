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
		cs := &card.CardSet{}
		c := &card.Card{1, 3}

		convey.Convey("push", func() {
			cs.PushCard(c)
			convey.So(cs.Cards, convey.ShouldContain, c)
		})
		convey.Convey("pop", func() {
			cs.PushCard(c)
			cp, _ := cs.PopCard()
			convey.So(cp, convey.ShouldEqual, c)
		})
		convey.Convey("sort", func() {
			c1 := &card.Card{4, 3}
			c2 := &card.Card{1, 3}
			c3 := &card.Card{2, 3}
			c4 := &card.Card{2, 2}

			cs.PushCard(c1)
			cs.PushCard(c4)
			cs.PushCard(c2)
			cs.PushCard(c3)
			cs = cs.Sort()

			csDst := &card.CardSet{}
			csDst.PushCard(c2)
			csDst.PushCard(c4)
			csDst.PushCard(c3)
			csDst.PushCard(c1)

			convey.So(cs, convey.ShouldResemble, csDst)
		})
	})
}
