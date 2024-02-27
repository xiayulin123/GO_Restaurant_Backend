package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/xiayulin123/GO_Restaurant_Backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	
	incomingRoutes.GET("/user", controller.GetUsers())
	incomingRoutes.GET("/user/:user_id", controller.GetUser())
	incomingRoutes.POST("/user/signup", controller.SignUp())
	incomingRoutes.POST("/user/login", controller.Login())


}