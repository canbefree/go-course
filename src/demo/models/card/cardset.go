package card

import (
	"errors"
)

type ICardSet interface {
	InsertCard(Card, int) *ICardSet
	InsertCards(cardsType) *ICardSet
	Sort() *ICardSet
	DelCard(Card) *ICardSet
	DelCards(cardsType) *ICardSet
	PushCard(cardsType) *ICardSet
	PushCards(cardsType) *ICardSet
	PopCard(cardsType) Card
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

func (cardSet *CardSet) InsertCards(cards cardsType) *CardSet {
	return cardSet
}

func (cardSet *CardSet) PushCard(card *Card) *CardSet {
	cardSet.Cards = append(cardSet.Cards, card)
	return cardSet
}

func (cardSet *CardSet) PopCard() (*Card, error) {
	len := len(cardSet.Cards)
	if len < 1 {
		return nil, errors.New("数组过短")
	}
	card := cardSet.Cards[len-1]
	cardSet.Cards = cardSet.Cards[0 : len-1]
	return card, nil
}

func (cardSet *CardSet) Sort() *CardSet {
	//冒泡排序
	for i := range cardSet.Cards {
		for j := 0; j < i; j++ {
			if cardSet.Cards[j].GetValue() > cardSet.Cards[j+1].GetValue() {
				cardSet.Cards[j], cardSet.Cards[j+1] = cardSet.Cards[j+1], cardSet.Cards[j]
			}
		}
	}
	return cardSet
}
