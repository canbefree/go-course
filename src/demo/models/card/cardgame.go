package card

type CardGame struct {
	cardSet CardSet
	Pair    int
}

type ICardGame interface {
	ExtCards() //多余的牌
	Init()     //初始化
}
