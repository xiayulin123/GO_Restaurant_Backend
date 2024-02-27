package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/xiayulin123/GO_Restaurant_Backend/controllers"

)

func FoodRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:food_id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/food/:food_id", controller.UpdateFood())

}