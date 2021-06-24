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

	router.GET("/warehouses", warehouseHandler.Get())

	router.Run()
}
