package main

import (
	"os"
	"github.com/gin-gonic/gin"

	"github.com/xiayulin123/GO_Restaurant_Backend/database"
	"github.com/xiayulin123/GO_Restaurant_Backend/routes"
	"github.com/xiayulin123/GO_Restaurant_Backend/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)	

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main () {
	port := os.Getenv("PORT")

	if port =="" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)


}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)

	return collection
}