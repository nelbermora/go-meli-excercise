package handler

import (
	"context"
	"time"

	productbatch "github.com/BenjaminBergerM/go-meli-exercise/internal/product_batch"
	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductBatch struct {
	productBatchService productbatch.Service
}

func NewProductBatch(s productbatch.Service) *ProductBatch {
	return &ProductBatch{
		productBatchService: s,
	}
}

func (l *ProductBatch) Store() gin.HandlerFunc {
	type request struct {
		Batch     string  `json:"batch_number"`
		Qty       int     `json:"current_quantity"`
		Temp      float32 `json:"current_temperature"`
		Due       string  `json:"due_date"`
		InitQty   int     `json:"initial_quantity"`
		ManufDate string  `json:"manufacturing_date"`
		MinTemp   float32 `json:"minumum_temperature"`
		Product   int     `json:"product_id"`
		Section   int     `json:"section_id"`
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

		if req.Batch == "" {
			c.JSON(422, web.NewError(422, "locality_name can not be empty"))
			return
		}

		ctx := context.Background()
		layout := "2006-01-02T15:04"
		dueDate, _ := time.Parse(layout, req.Due)
		manufDate, err := time.Parse(layout, req.ManufDate)
		if err != nil {
			switch err {

			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}
		emp, err := l.productBatchService.Store(ctx, req.Qty, req.InitQty, req.Product, req.Section, req.Temp, req.MinTemp, dueDate, manufDate, req.Batch)
		if err != nil {
			switch err {

			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}
		c.JSON(201, &response{emp})
	}
}
