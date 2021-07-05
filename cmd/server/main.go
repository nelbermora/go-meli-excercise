package main

import (
	"database/sql"

	"github.com/BenjaminBergerM/go-meli-exercise/cmd/server/handler"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/buyer"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/employee"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/product"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/seller"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/warehouse"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "./meli.db")
	router := gin.Default()

	warehouseService := warehouse.NewService()
	warehouseHandler := handler.NewWarehouse(warehouseService)
	warehousesRoutes := router.Group("/warehouses")
	{
		warehousesRoutes.GET("/", warehouseHandler.GetAll())
		warehousesRoutes.GET("/:id", warehouseHandler.Get())
		warehousesRoutes.POST("/", warehouseHandler.Store())
		warehousesRoutes.PATCH("/:id", warehouseHandler.Update())
		warehousesRoutes.DELETE("/:id", warehouseHandler.Delete())
	}

	sellerRepository := seller.NewRepository(db)
	sellerService := seller.NewService(sellerRepository)
	sellerHandler := handler.NewSeller(sellerService)
	sellersRoutes := router.Group("/sellers")
	{
		sellersRoutes.GET("/", sellerHandler.GetAll())
		sellersRoutes.GET("/:id", sellerHandler.Get())
		sellersRoutes.POST("/", sellerHandler.Store())
		sellersRoutes.PATCH("/:id", sellerHandler.Update())
		sellersRoutes.DELETE("/:id", sellerHandler.Delete())
	}

	buyerService := buyer.NewService()
	buyerHandler := handler.NewBuyer(buyerService)
	buyersRoutes := router.Group("/buyers")
	{
		buyersRoutes.GET("/", buyerHandler.GetAll())
		buyersRoutes.GET("/:id", buyerHandler.Get())
		buyersRoutes.POST("/", buyerHandler.Store())
		buyersRoutes.PATCH("/:id", buyerHandler.Update())
		buyersRoutes.DELETE("/:id", buyerHandler.Delete())
	}

	employeeService := employee.NewService()
	employeeHandler := handler.NewEmployee(employeeService)
	employeesRoutes := router.Group("/employees")
	{
		employeesRoutes.GET("/", employeeHandler.GetAll())
		employeesRoutes.GET("/:id", employeeHandler.Get())
		employeesRoutes.POST("/", employeeHandler.Store())
		employeesRoutes.PATCH("/:id", employeeHandler.Update())
		employeesRoutes.DELETE("/:id", employeeHandler.Delete())

	}

	productRepo := product.NewRepository(db)
	productService := product.NewService(productRepo)
	productHandler := handler.NewProduct(productService)
	productRoutes := router.Group("/products")
	{
		productRoutes.GET("/", productHandler.GetAll())
		productRoutes.GET("/:id", productHandler.Get())
		productRoutes.POST("/", productHandler.Store())
		productRoutes.PUT("/:id", productHandler.Update())
		productRoutes.DELETE("/:id", productHandler.Delete())
	}

	router.Run()
}
