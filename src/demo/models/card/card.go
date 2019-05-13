package card

//定义一个统一的类型 方便后面修改
type cardsType []*Card

type Card struct {
	Index int //大小
	Shape int //花色
}

func (card *Card) GetValue() int {
	//默认 主 Index 副 花色
	return card.Index*100 + card.Shape
}

type ICard interface {
	GetValue() int
}

func CompareCardValue(card1 Card, card2 Card) bool {
	return card1.GetValue() > card2.GetValue()
}

// HandCard 手牌
type HandCard struct {
	Cards cardsType
}

type PokerSet struct {
}
