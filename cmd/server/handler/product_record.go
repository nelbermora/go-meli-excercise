package handler

import (
	"context"
	"time"

	productrecord "github.com/BenjaminBergerM/go-meli-exercise/internal/product_record"
	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type Record struct {
	service productrecord.Service
}

func NewProductRecord(s productrecord.Service) *Record {
	return &Record{
		service: s,
	}
}
func (l *Record) Store() gin.HandlerFunc {
	type request struct {
		Update   string  `json:"last_update_date"`
		Purchase float32 `json:"purchase_price"`
		Sale     float32 `json:"sale_price"`
		Product  int     `json:"product_id"`
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

		if req.Product == 0 {
			c.JSON(422, web.NewError(422, "product_id can not be empty"))
			return
		}

		ctx := context.Background()
		layout := "2006-01-02"
		updateDate, err := time.Parse(layout, req.Update)
		if err != nil {
			switch err {

			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}
		resp, err := l.service.Store(ctx, updateDate, req.Purchase, req.Sale, req.Product)

		if err != nil {
			switch err {
			case productrecord.ErrProdExistance:
				c.JSON(409, web.NewError(409, err.Error()))
			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}
		c.JSON(201, &response{resp})
	}
}
