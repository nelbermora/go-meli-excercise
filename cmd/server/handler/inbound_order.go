package handler

import (
	"context"
	inboundorder "github.com/BenjaminBergerM/go-meli-exercise/internal/inbound_order"
	"time"

	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type InboundOrder struct {
	service inboundorder.Service
}

func NewInboundOrder(s inboundorder.Service) *InboundOrder {
	return &InboundOrder{
		service: s,
	}
}
func (l *InboundOrder) Store() gin.HandlerFunc {
	type request struct {
		OrderDate      string `json:"order_date"`
		OrderNumber    string `json:"order_number"`
		EmployeeId     int    `json:"employee_id"`
		ProductBatchId int    `json:"product_batch_id"`
		WarehouseId    int    `json:"warehouse_id"`
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

		ctx := context.Background()
		layout := "2006-01-02"
		orderDate, err := time.Parse(layout, req.OrderDate)
		if err != nil {
			switch err {

			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}
		resp, err := l.service.Store(ctx, orderDate, req.OrderNumber, req.EmployeeId, req.ProductBatchId, req.WarehouseId)

		if err != nil {
			switch err {
			case inboundorder.ErrEmployeeExistance:
				c.JSON(409, web.NewError(409, err.Error()))
			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}
		c.JSON(201, &response{resp})
	}
}
