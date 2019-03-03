package routers

import (
	"report/controllers"

	"github.com/gin-gonic/gin"
)

var Routers *gin.Engine

func init() {
	Routers = gin.Default()
}

func SetupRouters() {
	//real time,.. websocket
	//
	Routers.GET("/b/cards/save/:id_board", controllers.SaveCardsOnDB) // Bước tiền đề xây dựng database trước khi sử dụng /b/card/changeddue/:id
	Routers.GET("/b/cards/update/:id_board", controllers.UpdateCards)
	Routers.GET("/b/cards/review/:id_board", controllers.GetAllCardReview)
	Routers.GET("/b/cards/changedue/:id_board", controllers.GetAllCardChangeDue)
}
