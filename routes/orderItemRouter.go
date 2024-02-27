package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/xiayulin123/GO_Restaurant_Backend/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controller.GetOrderItems())
	incomingRoutes.GET("/orderItems/:orderItem_id", controller.GetOrderItem())
	incomingRoutes.POST("/orderItems", controller.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:orderItem_id", controller.UpdateOrderItem())

	incomingRoutes.GET("/orderItems-order/:order_id", controller.GetOrderItemsByOder())

}