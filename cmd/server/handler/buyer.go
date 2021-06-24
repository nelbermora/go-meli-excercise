package handler

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/buyer"
	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type Buyer struct {
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

func (w *Buyer) Get() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}

func (w *Buyer) Store() gin.HandlerFunc {
	type request struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
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
		if req.FirstName == "" {
			c.JSON(422, web.NewError(422, "first_name can not be empty"))
			return
		}
		if req.LastName == "" {
			c.JSON(422, web.NewError(422, "last_name can not be empty"))
			return
		}

		ctx := context.Background()
		buyer, err := w.buyerService.Store(ctx, req.FirstName, req.LastName)
		if err != nil {

		}

		c.JSON(201, &response{buyer})
	}
}

func (w *Buyer) Update() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}

func (w *Buyer) Delete() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}
