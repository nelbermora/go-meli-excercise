package main

import (
	"github.com/BenjaminBergerM/go-meli-exercise/cmd/server/handler"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/buyer"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/employee"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/warehouse"
	"github.com/gin-gonic/gin"
)

func main() {

	warehouseService := warehouse.NewService()
	warehouseHandler := handler.NewWarehouse(warehouseService)

	employeeService := employee.NewService()
	employeeHandler := handler.NewEmployee(employeeService)

	buyerService := buyer.NewService()
	buyerHandler := handler.NewBuyer(buyerService)

	router := gin.Default()

	warehousesRoutes := router.Group("/warehouses")
	{
		warehousesRoutes.GET("/", warehouseHandler.Get())
		warehousesRoutes.POST("/", warehouseHandler.Store())
		warehousesRoutes.PATCH("/:id", warehouseHandler.Get())
		warehousesRoutes.DELETE("/:id", warehouseHandler.Get())
	}
	buyersRoutes := router.Group("/buyers")
	{
		buyersRoutes.GET("/", buyerHandler.Get())
		buyersRoutes.POST("/", buyerHandler.Store())
		buyersRoutes.PATCH("/:id", buyerHandler.Get())
		buyersRoutes.DELETE("/:id", buyerHandler.Get())
	}
	employeesRoutes := router.Group("/employees")
	{
		employeesRoutes.GET("/", employeeHandler.Get())
		employeesRoutes.POST("/", employeeHandler.Store())
		employeesRoutes.PATCH("/:id", employeeHandler.Get())
		employeesRoutes.DELETE("/:id", employeeHandler.Get())
	}

	router.Run()
}
