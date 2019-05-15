package card

//一副牌
type CardPool struct {
	CardSet
}

type ICardPool interface {
	Init() //初始化手牌
}

func (cardPool *CardPool) Init() {
	// 3,4,5,6,7,8,9,10,j,Q,K,A,2
	indexs := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	shapes := [...]int{1, 2, 3, 4}

	cardSet := &CardSet{}
	for index := range indexs {
		for shape := range shapes {
			cardSet.PushCard(&Card{Index: index, Shape: shape})
		}
	}

	//大王小王
	cardSet.PushCard(&Card{Index: 14, Shape: 1})
	cardSet.PushCard(&Card{Index: 13, Shape: 2})

	cardPool.CardSet = *cardSet
}

func NewCardPool() *CardPool {
	return &CardPool{}
}
