package main

import (
	"github.com/BenjaminBergerM/go-meli-exercise/cmd/server/handler"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/warehouse"
	"github.com/gin-gonic/gin"
)

func main() {

	warehouseService := warehouse.NewService()
	warehouseHandler := handler.NewWarehouse(warehouseService)

	router := gin.Default()

	warehousesRoutes := router.Group("/warehouses")
	{
		warehousesRoutes.GET("/", warehouseHandler.Get())
		warehousesRoutes.POST("/", warehouseHandler.Store())
		warehousesRoutes.PATCH("/:id", warehouseHandler.Get())
		warehousesRoutes.DELETE("/:id", warehouseHandler.Get())
	}

	router.Run()
}
