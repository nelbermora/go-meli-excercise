package handler

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/buyer"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/employee"
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

func (b *Buyer) Get() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		paramID := c.Param("id")
		if paramID == "" {
			c.JSON(404, web.NewError(404, "not found"))
			return
		}

		ctx := context.Background()
		buyer, err := b.buyerService.Get(ctx, paramID)
		if err != nil {
			c.JSON(404, web.NewError(404, "buyer not found"))
			return
		}

		c.JSON(200, &response{buyer})
	}
}

func (b *Buyer) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Buyer `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		buyers, err := b.buyerService.GetAll(ctx)
		if err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		}

		c.JSON(200, &response{buyers})
	}
}

func (b *Buyer) Store() gin.HandlerFunc {
	type request struct {
		CardNumberID string `json:"card_number_id"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
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
		if req.CardNumberID == "" {
			c.JSON(422, web.NewError(422, "card_number_id can not be empty"))
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
		emp, err := b.buyerService.Store(ctx, req.CardNumberID, req.FirstName, req.FirstName)
		if err != nil {
			switch err {
			case employee.UNIQUE:
				c.JSON(409, web.NewError(409, err.Error()))
			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}

		c.JSON(201, &response{emp})
	}
}

func (b *Buyer) Update() gin.HandlerFunc {
	type request struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

		paramID := c.Param("id")
		if paramID == "" {
			c.JSON(404, web.NewError(404, "not found"))
			return
		}

		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(422, web.NewError(400, "json decoding: "+err.Error()))
			return
		}

		ctx := context.Background()
		emp, err := b.buyerService.Update(ctx, paramID, req.FirstName, req.LastName)
		if err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}

		c.JSON(200, &response{emp})
	}
}

func (b *Buyer) Delete() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		paramID := c.Param("id")
		if paramID == "" {
			c.JSON(404, web.NewError(404, "not found"))
			return
		}

		ctx := context.Background()
		err := b.buyerService.Delete(ctx, paramID)
		if err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}

		c.JSON(200, web.NewError(200, "The employee has been deleted"))
	}
}
