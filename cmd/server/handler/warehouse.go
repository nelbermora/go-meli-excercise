package handler

import (
	"context"
	"strconv"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
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
	type response struct {
		Data domain.Warehouse `json:"data"`
	}

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid ID"))
			return
		}
		ctx := context.Background()
		sel, err := w.warehouseService.Get(ctx, int(id))
		if err != nil {
			c.JSON(404, web.NewError(404, "warehouse not found"))
			return
		}
		c.JSON(201, &response{sel})
	}
}

func (w *Warehouse) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Warehouse `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		warehouses, err := w.warehouseService.GetAll(ctx)
		if err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		}

		c.JSON(201, &response{warehouses})
	}
}

func (w *Warehouse) Store() gin.HandlerFunc {
	type request struct {
		Address            string `json:"address"`
		Telephone          string `json:"telephone"`
		WarehouseCode      string `json:"warehouse_code"`
		MinimunCapacity    int    `json:"minimun_capacity"`
		MinimunTemperature int    `json:"minimun_temperature"`
		SectionNumber      int    `json:"section_number"`
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
		if req.MinimunCapacity == 0 {
			c.JSON(422, web.NewError(422, "minimun_capacity can not be empty"))
			return
		}
		if req.MinimunTemperature == 0 {
			c.JSON(422, web.NewError(422, "minimun_temperature can not be empty"))
			return
		}
		if req.SectionNumber == 0 {
			c.JSON(422, web.NewError(422, "section_number can not be empty"))
			return
		}

		ctx := context.Background()
		wh, err := w.warehouseService.Store(ctx, req.Address, req.Telephone, req.WarehouseCode, req.MinimunCapacity, req.MinimunTemperature, req.SectionNumber)
		if err != nil {
			switch err {
			case warehouse.UNIQUE:
				c.JSON(409, web.NewError(409, err.Error()))
			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}

		c.JSON(201, &response{wh})
	}
}

func (w *Warehouse) Update() gin.HandlerFunc {
	type request struct {
		Address            string `json:"address"`
		Telephone          string `json:"telephone"`
		WarehouseCode      string `json:"warehouse_code"`
		MinimunCapacity    int    `json:"minimun_capacity"`
		MinimunTemperature int    `json:"minimun_temperature"`
		SectionNumber      int    `json:"section_number"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid ID"))
			return
		}

		if err := c.Bind(&req); err != nil {
			c.JSON(422, web.NewError(422, "json decoding: "+err.Error()))
			return
		}

		ctx := context.Background()
		warehouse, err := w.warehouseService.Update(ctx, int(id), req.Address, req.Telephone, req.WarehouseCode, req.MinimunCapacity, req.MinimunTemperature, req.SectionNumber)
		if err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}

		c.JSON(200, &response{warehouse})
	}
}

func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid ID"))
			return
		}

		ctx := context.Background()
		err = w.warehouseService.Delete(ctx, int(id))
		if err != nil {
			c.JSON(400, web.NewError(400, err.Error()))
			return
		}

		c.JSON(200, web.NewError(200, "The warehouse has been deleted"))
	}
}
