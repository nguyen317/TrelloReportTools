package controllers

import (
	"github.com/adlio/trello"
)

func GetCardsIsOpenOnWeek(idBoard, nameList string) ([]*trello.Card, error) {
	cards, err := client.SearchCards("board:"+idBoard+" is:open sort:created created:week list:"+nameList, trello.Defaults())
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func GetCardsOnBoard(idBoard string) ([]*trello.Card, error) {
	board, err := client.GetBoard(idBoard, trello.Defaults())
	if err != nil {
		return nil, err
	}

	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		return nil, err
	}

	return cards, nil
}
