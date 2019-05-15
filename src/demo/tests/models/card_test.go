package models

import (
	"demo/models/card"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCard(t *testing.T) {
	Convey("testCard", t, func() {
		card := &card.Card{Index: 12, Shape: 3}
		Convey("testCardGetValue", func() {
			So(card.GetValue(), ShouldEqual, 1203)
		})
	})
}

func TestCardSet(t *testing.T) {
	Convey("CardSet", t, func() {
		c1 := &card.Card{1, 2}
		c2 := &card.Card{1, 4}
		c3 := &card.Card{2, 3}
		c4 := &card.Card{5, 2}
		object := &card.CardSet{Cards: []*card.Card{c1, c4, c2, c3}}
		c := &card.Card{1, 3}
		Convey("rand", func() {
			subject := object
			except := object.Rand()
			So(subject.Cards, ShouldResemble, except.Cards)

		})
		Convey("del-global", func() {
			object.PushCard(c2)
			object.DelCard(c2, true)
			except := &card.CardSet{Cards: []*card.Card{c1, c4, c3}}
			So(object, ShouldResemble, except)
		})

		Convey("del-one", func() {
			object.PushCard(c2)
			object.DelCard(c2, false)
			except := &card.CardSet{Cards: []*card.Card{c1, c4, c3, c2}}
			So(object, ShouldResemble, except)
		})

		Convey("unset", func() {
			object.Unset(1)
			except := &card.CardSet{Cards: []*card.Card{c1, c2, c3}}
			So(object, ShouldResemble, except)
		})

		Convey("inserts", func() {
			object.InsertCards([]*card.Card{c1, c2})
			except := &card.CardSet{Cards: []*card.Card{c1, c4, c2, c3, c1, c2}}
			So(object, ShouldResemble, except)
		})

		Convey("insert", func() {
			object.InsertCard(c, 3)
			except := &card.CardSet{Cards: []*card.Card{c1, c4, c2, c, c3}}
			So(object, ShouldResemble, except)
		})
		Convey("push", func() {
			object.PushCard(c)
			So(object.Cards, ShouldContain, c)
		})
		Convey("pop", func() {
			cp, _ := object.PopCard()
			So(cp, ShouldEqual, c3)
		})
		Convey("sort", func() {
			object = object.Sort()
			except := &card.CardSet{Cards: []*card.Card{c1, c2, c3, c4}}
			So(object, ShouldResemble, except)
		})
	})
}

func TestCardPool(t *testing.T) {
	Convey("test pool", t, func() {
		Convey("init", func() {
			cardPool := card.NewCardPool()
			cardPool.Init()
			//一副牌应该是54张
			So(len(cardPool.Cards), ShouldEqual, 54)
			So(cardPool.Cards, ShouldContain, &card.Card{Index: 9, Shape: 3})
		})
	})
}
