package main

import (
	"database/sql"
	"github.com/BenjaminBergerM/go-meli-exercise/cmd/server/handler"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/buyer"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/employee"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/product"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/warehouse"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
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
		warehousesRoutes.GET("/", warehouseHandler.GetAll())
		warehousesRoutes.GET("/:id", warehouseHandler.Get())
		warehousesRoutes.POST("/", warehouseHandler.Store())
		warehousesRoutes.PATCH("/:id", warehouseHandler.Update())
		warehousesRoutes.DELETE("/:id", warehouseHandler.Delete())
	}
	buyersRoutes := router.Group("/buyers")
	{
		buyersRoutes.GET("/", buyerHandler.GetAll())
		buyersRoutes.GET("/:id", buyerHandler.Get())
		buyersRoutes.POST("/", buyerHandler.Store())
		buyersRoutes.PATCH("/:id", buyerHandler.Update())
		buyersRoutes.DELETE("/:id", buyerHandler.Delete())
	}
	employeesRoutes := router.Group("/employees")
	{
		employeesRoutes.GET("/", employeeHandler.GetAll())
		employeesRoutes.GET("/:id", employeeHandler.Get())
		employeesRoutes.POST("/", employeeHandler.Store())
		employeesRoutes.PATCH("/:id", employeeHandler.Update())
		employeesRoutes.DELETE("/:id", employeeHandler.Delete())

	}
	db, _ := sql.Open("sqlite3", "./meli.db")
	productRepo := product.NewRepository(db)
	productService := product.NewService(productRepo)
	productHandler := handler.NewProduct(productService)
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("/", productHandler.GetAll())
		productRoutes.POST("/", productHandler.Store())
		productRoutes.PATCH("/:id", productHandler.Get())
		productRoutes.DELETE("/:id", productHandler.Get())
	}

	router.Run()
}
