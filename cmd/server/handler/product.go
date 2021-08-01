package handler

import (
	"context"
	"strconv"

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

func (p *Product) GetAll() gin.HandlerFunc {

	type response struct {
		Data []domain.Product `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		products, err := p.productService.GetAll(ctx)
		if err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		}

		c.JSON(200, &response{products})
	}
}

func (p *Product) Get() gin.HandlerFunc {

	type response struct {
		Data domain.Product `json:"data"`
	}

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid ID"))
			return
		}
		ctx := context.Background()
		prod, err := p.productService.Get(ctx, int(id))
		if err != nil {
			c.JSON(404, web.NewError(404, "Product not found"))
			return
		}

		c.JSON(200, &response{prod})
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
		ProductCode    string  `json:"product_code"`
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
		prod, err := p.productService.Store(ctx, req.Description, req.ProductCode, req.Height, req.Length, req.Netweight, req.RecomFreezTemp, req.Width, req.ProductTypeID, req.SellerID, req.ExpirationRate, req.FreezingRate)
		if err != nil {
			switch err {
			case product.UNIQUE:
				c.JSON(409, web.NewError(409, err.Error()))
			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}

		c.JSON(201, &response{prod})
	}
}

func (p *Product) Update() gin.HandlerFunc {

	type request struct {
		Description    string  `json:"description"`
		ExpirationRate int     `json:"expiration_rate"`
		FreezingRate   int     `json:"freezing_rate"`
		Height         float32 `json:"height"`
		Length         float32 `json:"length"`
		Netweight      float32 `json:"netweight"`
		ProductCode    string  `json:"product_code"`
		RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
		Width          float32 `json:"width"`
		ProductTypeID  int     `json:"product_type_id"`
		SellerID       int     `json:"seller_id"`
	}

	type response struct {
		Data string `json:"data"`
	}

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid ID"))
			return
		}

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
		prod, err := p.productService.Update(ctx, int(id), req.Description, req.ProductCode, req.Height, req.Length, req.Netweight, req.RecomFreezTemp, req.Width, req.ProductTypeID, req.SellerID, req.ExpirationRate, req.FreezingRate)
		if err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}

		c.JSON(200, prod)
	}
}

func (p *Product) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid ID"))
			return
		}

		ctx := context.Background()
		err = p.productService.Delete(ctx, int(id))
		if err != nil {
			c.JSON(400, web.NewError(400, err.Error()))
			return
		}

		c.JSON(200, web.NewError(200, "The product has been deleted"))
	}
}

func (p *Product) GetProductsByRecord() gin.HandlerFunc {
	type response struct {
		Data []domain.Product `json:"data"`
	}
	return func(c *gin.Context) {
		id := c.Query("id")
		intId, _ := strconv.Atoi(id)
		ctx := context.Background()
		rep, err := p.productService.GetByRecord(ctx, intId)
		if err != nil {
			switch err {
			case product.ErrProdExistance:
				c.JSON(404, web.NewError(404, "product not found"))
				return
			default:
				c.JSON(500, web.NewError(500, err.Error()))
				return
			}
		}
		c.JSON(200, &response{rep})
	}
}
