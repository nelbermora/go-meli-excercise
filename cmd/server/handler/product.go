package handler

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/product"
	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	productService product.Service
}

func NewProduct(w product.Service) *Product {
	return &Product{
		productService: w,
	}
}

func (w *Product) Get() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}

func (p *Product) Store() gin.HandlerFunc {
	type request struct {
		Description    string  `json:"description"`
		ExpirationRate int     `json:"expiration_rate"`
		FreezingRate   int     `json:"freezing_rate"`
		Height         float32 `json:"height"`
		Length         float32 `json:"length"`
		Netweight      float32 `json:"netweight"`
		ProductCode    string     `json:"product_code"`
		RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
		Width          float32 `json:"width"`
		ProductTypeID  int     `json:"product_type_id"`
		SellerID       int     `json:"seller_id"`
	}

	type response struct {
		Data domain.Product `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		if err := c.Bind(&req); err != nil {
			c.JSON(422, web.NewError(400, "json decoding: "+err.Error()))
			return
		}
		if req.Height == 0 {
			c.JSON(422, web.NewError(422, "height can not be empty"))
			return
		}
		if req.Length == 0 {
			c.JSON(422, web.NewError(422, "length can not be empty"))
			return
		}
		if req.Netweight == 0 {
			c.JSON(422, web.NewError(422, "netweight can not be empty"))
			return
		}

		if req.Width == 0 {
			c.JSON(422, web.NewError(422, "width can not be empty"))
			return
		}
		if req.ProductCode == "" {
			c.JSON(422, web.NewError(422, "product_code can not be empty"))
			return
		}

		ctx := context.Background()
		warehouse, err := p.productService.Store(ctx, req.Description, req.ProductCode, req.Height, req.Length, req.Netweight, req.RecomFreezTemp, req.Width, req.ProductTypeID, req.SellerID, req.ExpirationRate, req.FreezingRate)
		if err != nil {

		}

		c.JSON(201, &response{warehouse})
	}
}

func (w *Product) Update() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}

func (w *Product) Delete() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}
