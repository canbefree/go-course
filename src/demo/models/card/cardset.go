package card

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

//牌堆
type ICardSet interface {
	InsertCard(Card, int) *ICardSet
	InsertCards(cardsType) *ICardSet
	Sort() *ICardSet
	DelCard(Card, bool) *ICardSet //删除 bool 为true 删除所有的。 false只删除第一个找到的
	DelCards(cardsType) *ICardSet
	Unset(int) *ICardSet
	PushCard(cardsType) *ICardSet
	PushCards(cardsType) *ICardSet
	PopCard(cardsType) Card
	PopCards(cardsType) *ICardSet
	Rand() *ICardSet //洗牌
}

type CardSet struct {
	Cards cardsType
}

func (cardSet *CardSet) InsertCard(card *Card, i int) *CardSet {
	if i > len(cardSet.Cards) {
		fmt.Printf("插入失败 %v %v大于长度", cardSet, i)
		return cardSet
	}
	j := len(cardSet.Cards)
	cardSet.Cards = append(cardSet.Cards, nil)
	for {
		if j > i {
			cardSet.Cards[j] = cardSet.Cards[j-1]
			j--
		} else {
			break
		}
	}
	cardSet.Cards[i] = card
	return cardSet
}
func (cardSet *CardSet) DelCard(card *Card, global bool) *CardSet {
	for k, item := range cardSet.Cards {
		if card.GetValue() == item.GetValue() {
			cardSet = cardSet.Unset(k)
			if !global {
				break
			}
		}
	}
	return cardSet
}

func (cardSet *CardSet) Unset(i int) *CardSet {
	if isOverRange(cardSet.Cards, i) {
		return cardSet
	}
	j := i
	for {
		if j < len(cardSet.Cards)-1 {
			cardSet.Cards[j] = cardSet.Cards[j+1]
			j++
		} else {
			break
		}
	}
	cardSet.PopCard()
	return cardSet
}

func (cardSet *CardSet) DelCards(cardsType) *CardSet {
	return cardSet
}

func (cardSet *CardSet) InsertCards(cards cardsType) *CardSet {
	cardSet.Cards = append(cardSet.Cards, cards...)
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

func (cardSet *CardSet) Rand() *CardSet {
	rand.Seed(time.Now().UnixNano())

	for i := len(cardSet.Cards) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		cardSet.Cards[i], cardSet.Cards[num] = cardSet.Cards[num], cardSet.Cards[i]
	}

	return cardSet
}

func isOverRange(item cardsType, i int) bool {
	if i < 0 || i > len(item) {
		fmt.Printf("cardSet over range!")
		return true
	}
	return false
}
