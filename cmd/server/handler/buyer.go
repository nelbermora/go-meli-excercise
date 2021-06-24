package handler

import (
	"context"
	"strconv"

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

func (b *Buyer) Get() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.JSON(422, web.NewError(422, "id must be an integer"))
			return
		}

		ctx := context.Background()
		employee, err := b.buyerService.Get(ctx, id)
		if err != nil {

		}

		c.JSON(200, &response{employee})
	}
}

func (b *Buyer) GetAll() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		ctx := context.Background()
		employee, err := b.buyerService.GetAll(ctx)
		if err != nil {

		}
		c.JSON(201, &response{employee})
	}
}

func (b *Buyer) Store() gin.HandlerFunc {
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
		buyer, err := b.buyerService.Store(ctx, req.FirstName, req.LastName)
		if err != nil {

		}

		c.JSON(201, &response{buyer})
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
		var req request

		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.JSON(422, web.NewError(422, "id must be an integer"))
			return
		}
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
		buyer, err := b.buyerService.Update(ctx, id, req.FirstName, req.LastName)
		if err != nil {

		}

		c.JSON(201, &response{buyer})
	}
}

func (b *Buyer) Delete() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		paramID := c.Param("id")
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.JSON(422, web.NewError(422, "id must be an integer"))
			return
		}

		ctx := context.Background()
		employee, err := b.buyerService.Delete(ctx, id)
		if err != nil {

		}

		c.JSON(201, &response{employee})
	}
}
