package handler

import (
	"context"
	purchaseorder "github.com/BenjaminBergerM/go-meli-exercise/internal/purchase_order"
	"time"

	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type PurchaseOrder struct {
	service purchaseorder.Service
}

func NewPurchaseOrder(s purchaseorder.Service) *PurchaseOrder {
	return &PurchaseOrder{
		service: s,
	}
}

func (l *PurchaseOrder) Store() gin.HandlerFunc {
	type request struct {
		OrderNumber     string `json:"order_number"`
		OrderDate       string `json:"order_date"`
		TrackingCode    string `json:"tracking_code"`
		BuyerId         int    `json:"buyer_id"`
		ProductRecordId int    `json:"product_record_id"`
		OrderStatusId   int    `json:"order_status_id"`
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
		resp, err := l.service.Store(ctx, req.OrderNumber, orderDate, req.TrackingCode, req.BuyerId, req.ProductRecordId, req.OrderStatusId)

		if err != nil {
			switch err {
			case purchaseorder.ErrBuyerExistance:
				c.JSON(409, web.NewError(409, err.Error()))
			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}
		c.JSON(201, &response{resp})
	}
}
