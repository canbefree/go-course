package card

//一副牌
type CardPool struct {
	cardSet CardSet
}

type ICardPool interface {
}

//洗牌
func (cardPool CardPool) Rand() {
	cardPool.cardSet.Rand()
}
