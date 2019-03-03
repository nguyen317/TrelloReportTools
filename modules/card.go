package modules

import (
	"time"

	"github.com/adlio/trello"
)

type Card struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	TimeGuessForDone int          `json:"timeGuessForDone"`
	TimeRealForDone  int          `json:"timeRealForDone"`
	IsEditedDueDate  bool         `json:"isEditedDueDate"`
	Due              *time.Time   `json:"due"`
	HistoryDueDate   []*time.Time `json:"historyDueDate"`
}

func (c Card) NewCard(card *trello.Card) Card {
	c.ID = card.ID
	c.Name = card.Name
	c.Due = card.Due
	return c
}
