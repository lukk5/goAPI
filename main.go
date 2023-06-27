package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goAPI/pkg/extensions"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	itemHandler := extensions.InitializeItemHandler()

	// Serve Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/item", itemHandler.GetItemsHandler)
	router.GET("/item/:id", itemHandler.GetItemByIdHandler)
	router.POST("/item/add", itemHandler.AddItemHandler)

	router.Run(":8080")
}
