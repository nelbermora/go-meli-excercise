package handler

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/warehouse"
	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	warehouseService warehouse.Service
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		warehouseService: w,
	}
}

func (w *Warehouse) Get() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}

func (w *Warehouse) Store() gin.HandlerFunc {
	type request struct {
		Address       string `json:"address"`
		Telephone     string `json:"telephone"`
		WarehouseCode string `json:"warehouse_code"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		if err := c.Bind(&req); err != nil {
			c.JSON(422, web.NewError(422, "json decoding: "+err.Error()))
			return
		}
		if req.Address == "" {
			c.JSON(422, web.NewError(422, "address can not be empty"))
			return
		}
		if req.Telephone == "" {
			c.JSON(422, web.NewError(422, "telephone can not be empty"))
			return
		}
		if req.WarehouseCode == "" {
			c.JSON(422, web.NewError(422, "warehouse_code can not be empty"))
			return
		}

		ctx := context.Background()
		warehouse, err := w.warehouseService.Store(ctx, req.Address, req.Telephone, req.WarehouseCode)
		if err != nil {

		}

		c.JSON(201, &response{warehouse})
	}
}

func (w *Warehouse) Update() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}

func (w *Warehouse) Delete() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}
