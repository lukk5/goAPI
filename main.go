package main

import (
	"goAPI/pkg/extensions"
	"goAPI/pkg/handlers"
	"goAPI/pkg/middlewares"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router      *gin.Engine
	ItemHandler *handlers.ItemHandler
}

func (a *App) Initialize() {
	a.Router = gin.Default()
	a.ItemHandler = extensions.InitializeItemHandler()

	a.initializeRoutes()
}

func (a *App) initializeRoutes() {

	// Serve Swagger documentation
	a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Apply exception handling middleware
	a.Router.Use(middlewares.ExceptionHandler)

	// Item routes
	a.Router.GET("/item", a.ItemHandler.GetItemsHandler)
	a.Router.GET("/item/:id", a.ItemHandler.GetItemByIdHandler)
	a.Router.POST("/item/add", a.ItemHandler.AddItemHandler)
}

func (a *App) Run(addr string) {
	a.Router.Run(addr)
}

func main() {
	app := &App{}
	app.Initialize()
	app.Run(":8080")
}
