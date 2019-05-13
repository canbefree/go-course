package card

type PokerGame struct {
}

func (pokergame *PokerGame) ParseCard(value int) *Card {
	return &Card{Index: 1}
}
