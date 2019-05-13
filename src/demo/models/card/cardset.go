package card

type ICardSet interface {
	InsertCard(Card, int) *ICardSet
	InsertCards(cardsType) *ICardSet
	Sort() *ICardSet
	DelCard(Card) *ICardSet
	DelCards(cardsType) *ICardSet
	PushCard(cardsType) *ICardSet
	PushCards(cardsType) *ICardSet
	PopCard(cardsType) *ICardSet
	PopCards(cardsType) *ICardSet
}

type CardSet struct {
	Cards cardsType
}

func (cardSet *CardSet) InsertCard(card *Card, i int) *CardSet {
	if i >= len(cardSet.Cards) {
	}
	cardSet.Cards[i] = card
	return cardSet
}

func (cardSet *CardSet) InsertCards(cards cardsType, i int) *CardSet {
	return cardSet
}
