package controllers

import (
	"TrelloReportTools/database"
	"TrelloReportTools/modules"

	"github.com/adlio/trello"
	"github.com/gin-gonic/gin"
)

func GetAllCardReview(c *gin.Context) {
	idBoard := c.Param("id_board")
	var tmpcard modules.Card
	var myCardsOnReviewMe []modules.Card
	var myCardsOnDone []modules.Card

	cardsOnListReviewMe, err := GetCardsIsOpenOnWeek(idBoard, "review-me")
	if err != nil {
		// Handle error
	}
	cardsOnListDone, err := GetCardsIsOpenOnWeek(idBoard, "Done")
	if err != nil {
		// Handle error
	}

	for _, v := range cardsOnListReviewMe {
		myCardsOnReviewMe = append(myCardsOnReviewMe, tmpcard.NewCard(v))
	}
	for _, v := range cardsOnListDone {
		myCardsOnDone = append(myCardsOnDone, tmpcard.NewCard(v))
	}

	c.JSON(200, gin.H{
		"List card on review-me": cardsOnListReviewMe,
		"List card on Done":      cardsOnListDone,
	})

}

func GetAllCardChangeDue(c *gin.Context) {
	idBoard := c.Param("id_board")
	cardsOnBoard, err := GetCardsOnBoard(idBoard)
	if err != nil {
		// Handle error
	}

	cardsOnDB := database.GetCards()
	var cardsChangedDueDate []*trello.Card
	for i, _ := range cardsOnBoard {
		for j, _ := range cardsOnDB {
			if CheckChangeDue(cardsOnBoard[i], cardsOnDB[j]) {
				cardsChangedDueDate = append(cardsChangedDueDate, cardsOnBoard[i])
			}
		}
	}

	c.JSON(200, gin.H{
		"Cards changed due date": cardsChangedDueDate,
	})
}

// Check due date of 2 card
func CheckChangeDue(cardsBoard *trello.Card, cardsDB modules.Card) bool {
	if cardsBoard.ID != cardsDB.ID {
		return false
	}
	if cardsBoard.Due == nil && cardsDB.Due == nil {
		return false
	}
	if cardsBoard.Due == nil && cardsDB.Due != nil {
		return true
	}
	if cardsBoard.Due != nil && cardsDB.Due == nil {
		return true
	}
	if cardsBoard.Due.String() == cardsDB.Due.String() {
		return false
	}
	return true
}

func SaveCardsOnDB(c *gin.Context) {
	idBoard := c.Param("id_board")
	cardsOnBoard, err := GetCardsOnBoard(idBoard)
	if err != nil {
		// Handle error
	}

	for _, v := range cardsOnBoard {
		value := *v
		tmpCard := modules.Card{
			ID:   value.ID,
			Name: value.Name,
			Due:  value.Due,
		}
		database.SaveCard(tmpCard)
	}

	cards := database.GetCards()
	c.JSON(200, gin.H{
		"Cards on database": cards,
	})
}

func UpdateCards(c *gin.Context) {
	idBoard := c.Param("id_board")
	cardsOnBoard, err := GetCardsOnBoard(idBoard)
	if err != nil {
		// Handle error
	}
	for _, v := range cardsOnBoard {
		tmpCard := modules.Card{
			ID:   v.ID,
			Name: v.Name,
			Due:  v.Due,
		}
		database.UpdateCard(tmpCard)
	}

	c.JSON(200, gin.H{
		"Update card": database.GetCards(),
	})
}
