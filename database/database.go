package database

import (
	"report/modules"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	OpenConnection()
	AutoMigrate()
	CloseConnection()
}
func GetCards() []modules.Card {
	OpenConnection()
	defer CloseConnection()

	var cards []modules.Card
	db.Find(&cards)

	return cards

}

func SaveCard(card modules.Card) {
	OpenConnection()
	defer CloseConnection()
	// fmt.Println()
	// if db.NewRecord(card) {
	// 	fmt.Println("create: ")
	// 	fmt.Println(card)
	// 	db.Create(&card)
	// }
	db.Create(&card)
}

func UpdateCard(card modules.Card) {
	OpenConnection()
	defer CloseConnection()
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

func OpenConnection() {
	var err error
	db, err = gorm.Open("sqlite3", "card.db")
	if err != nil {
		// Handle error
	}
}
func AutoMigrate() {
	db.AutoMigrate(&modules.Card{})
}
func CloseConnection() {
	db.Close()
}
