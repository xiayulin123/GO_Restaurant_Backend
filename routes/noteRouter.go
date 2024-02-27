package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/xiayulin123/GO_Restaurant_Backend/controllers"
)

func NoteRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/notes", controller.GetNotes())
	incomingRoutes.GET("/notes/:note_id", controller.GetNote())
	incomingRoutes.POST("/notes", controller.CreateNote())
	incomingRoutes.PATCH("/notes/:note_id", controller.UpdateNote())
}