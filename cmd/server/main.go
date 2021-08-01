package main

import (
	"database/sql"
	purchaseorder "github.com/BenjaminBergerM/go-meli-exercise/internal/purchase_order"

	"github.com/BenjaminBergerM/go-meli-exercise/cmd/server/handler"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/buyer"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/carry"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/employee"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/locality"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/product"
	productbatch "github.com/BenjaminBergerM/go-meli-exercise/internal/product_batch"
	productrecord "github.com/BenjaminBergerM/go-meli-exercise/internal/product_record"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/section"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/seller"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/warehouse"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, _ := sql.Open("sqlite3", "./meli.db")
	router := gin.Default()

	warehouseRepository := warehouse.NewRepository(db)
	warehouseService := warehouse.NewService(warehouseRepository)
	warehouseHandler := handler.NewWarehouse(warehouseService)
	warehousesRoutes := router.Group("/api/v1/warehouses")
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
	sellersRoutes := router.Group("/api/v1/sellers")
	{
		sellersRoutes.GET("/", sellerHandler.GetAll())
		sellersRoutes.GET("/:id", sellerHandler.Get())
		sellersRoutes.POST("/", sellerHandler.Store())
		sellersRoutes.PATCH("/:id", sellerHandler.Update())
		sellersRoutes.DELETE("/:id", sellerHandler.Delete())
	}

	sectionRepository := section.NewRepository(db)
	sectionService := section.NewService(sectionRepository)
	sectionHandler := handler.NewSection(sectionService)
	sectionsRoutes := router.Group("/api/v1/sections")
	{
		sectionsRoutes.GET("/", sectionHandler.GetAll())
		sectionsRoutes.GET("/:id", sectionHandler.Get())
		sectionsRoutes.GET("/reports", sectionHandler.GetBatchesBySection())
		sectionsRoutes.POST("/", sectionHandler.Store())
		sectionsRoutes.PATCH("/:id", sectionHandler.Update())
		sectionsRoutes.DELETE("/:id", sectionHandler.Delete())
	}

	productRepo := product.NewRepository(db)
	productService := product.NewService(productRepo)
	productHandler := handler.NewProduct(productService)
	productRoutes := router.Group("/api/v1/products")
	{
		productRoutes.GET("/", productHandler.GetAll())
		productRoutes.GET("/:id", productHandler.Get())
		productRoutes.POST("/", productHandler.Store())
		productRoutes.PUT("/:id", productHandler.Update())
		productRoutes.DELETE("/:id", productHandler.Delete())
		productRoutes.GET("/reportRecords", productHandler.GetProductsByRecord())
	}

	employeeRepository := employee.NewRepository(db)
	employeeService := employee.NewService(employeeRepository)
	employeeHandler := handler.NewEmployee(employeeService)
	employeesRoutes := router.Group("/api/v1/employees")
	{
		employeesRoutes.GET("/", employeeHandler.GetAll())
		employeesRoutes.GET("/:id", employeeHandler.Get())
		employeesRoutes.POST("/", employeeHandler.Store())
		employeesRoutes.PATCH("/:id", employeeHandler.Update())
		employeesRoutes.DELETE("/:id", employeeHandler.Delete())

	}

	buyerRepository := buyer.NewRepository(db)
	buyerService := buyer.NewService(buyerRepository)
	buyerHandler := handler.NewBuyer(buyerService)
	buyersRoutes := router.Group("/api/v1/buyers")
	{
		buyersRoutes.GET("/", buyerHandler.GetAll())
		buyersRoutes.GET("/:id", buyerHandler.Get())
		buyersRoutes.POST("/", buyerHandler.Store())
		buyersRoutes.PATCH("/:id", buyerHandler.Update())
		buyersRoutes.DELETE("/:id", buyerHandler.Delete())
	}

	localitiesRepository := locality.NewRepository(db)
	localitiesService := locality.NewService(localitiesRepository)
	localitiesHandler := handler.NewLocality(localitiesService)
	localitiesRoutes := router.Group("/api/v1/localities")
	{
		localitiesRoutes.POST("/", localitiesHandler.Store())
		localitiesRoutes.GET("/reports", localitiesHandler.GetSellersByLoc())
		localitiesRoutes.GET("/reportsCarries", localitiesHandler.GetCarriesByLoc())
	}

	carryRepository := carry.NewRepository(db)
	carryService := carry.NewService(carryRepository, localitiesRepository)
	carryHandler := handler.NewCarry(carryService)
	carryRoutes := router.Group("/api/v1/carries")
	{
		carryRoutes.POST("/", carryHandler.Store())
	}

	batchRepository := productbatch.NewRepository(db)
	batchSvc := productbatch.NewService(batchRepository, sectionRepository, productRepo)
	batchHandler := handler.NewProductBatch(batchSvc)
	batchRoutes := router.Group("/api/v1/productBatches")
	{
		batchRoutes.POST("/", batchHandler.Store())
	}

	recordRepo := productrecord.NewRepository(db)
	recordSvc := productrecord.NewService(recordRepo, productRepo)
	recordHandler := handler.NewProductRecord(recordSvc)
	recordRoutes := router.Group("/api/v1/productRecords")
	{
		recordRoutes.POST("/", recordHandler.Store())
	}

	purchaseRepo := purchaseorder.NewRepository(db)
	purchaseSvc := purchaseorder.NewService(purchaseRepo, buyerRepository)
	purchaseHandler := handler.NewPurchaseOrder(purchaseSvc)
	purchaseRoutes := router.Group("/api/v1/purchaseOrders")
	{
		purchaseRoutes.POST("/", purchaseHandler.Store())
	}


	router.Run()
}
