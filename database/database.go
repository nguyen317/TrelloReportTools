package database

import (
	"TrelloReportTools/modules"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open("sqlite3", "card.db")
	if err != nil {
		// Handle error
	}

	db.AutoMigrate(&modules.Card{})
}

func GetCards() []modules.Card {
	var cards []modules.Card
	db.Find(&cards)

	return cards
}

func SaveCard(card modules.Card) {
	db.Create(&card)
}

func UpdateCard(card modules.Card) {
	var newCard modules.Card
	db.Where("id = ?", card.ID).First(&newCard)

	if newCard.ID == "" { // Create new card
		db.Create(&card)
	} else { // Update old card
		newCard.Name = card.Name
		newCard.Due = card.Due
		db.Save(&newCard)
	}
}
