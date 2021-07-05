package handler

import (
	"context"
	"strconv"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/seller"
	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type Seller struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
	}
}

func (s *Seller) GetAll() gin.HandlerFunc {

	type response struct {
		Data []domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		sellers, err := s.sellerService.GetAll(ctx)
		if err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		}

		c.JSON(201, &response{sellers})
	}
}

func (s *Seller) Get() gin.HandlerFunc {

	type response struct {
		Data domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid ID"))
			return
		}
		ctx := context.Background()
		sel, err := s.sellerService.Get(ctx, int(id))
		if err != nil {
			c.JSON(404, web.NewError(404, "Seller not found"))
			return
		}
		c.JSON(201, &response{sel})
	}
}

func (s *Seller) Store() gin.HandlerFunc {
	type request struct {
		CID         int    `json:"cid"`
		CompanyName string `json:"company_name"`
		Address     string `json:"address"`
		Telephone   string `json:"telephone"`
		LocalityID  int    `json:"locality_id"`
	}

	type response struct {
		Data domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		if err := c.Bind(&req); err != nil {
			c.JSON(422, web.NewError(400, "json decoding: "+err.Error()))
			return
		}
		if req.CID == 0 {
			c.JSON(422, web.NewError(422, "cid can not be empty"))
			return
		}
		if req.CompanyName == "" {
			c.JSON(422, web.NewError(422, "company_name can not be empty"))
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
		if req.LocalityID == 0 {
			c.JSON(422, web.NewError(422, "locality_id can not be empty"))
			return
		}

		ctx := context.Background()
		sel, err := s.sellerService.Store(ctx, req.CID, req.CompanyName, req.Address, req.Telephone, req.LocalityID)
		if err != nil {
			switch err {
			case seller.UNIQUE:
				c.JSON(409, web.NewError(409, err.Error()))
			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}

		c.JSON(201, &response{sel})
	}
}

func (s *Seller) Update() gin.HandlerFunc {

	type request struct {
		SellerID    int    `json:"seller_id"`
		CID         int    `json:"cid"`
		CompanyName string `json:"company_name"`
		Address     string `json:"address"`
		Telephone   string `json:"telephone"`
		LocalityID  int    `json:"locality_id"`
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

		ctx := context.Background()
		sel, err := s.sellerService.Update(ctx, int(id), req.CID, req.CompanyName, req.Address, req.Telephone, req.LocalityID)
		if err != nil {
			c.JSON(400, web.NewError(400, err.Error()))
			return
		}

		c.JSON(200, sel)
	}
}

func (s *Seller) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid ID"))
			return
		}

		ctx := context.Background()
		err = s.sellerService.Delete(ctx, int(id))
		if err != nil {
			c.JSON(400, web.NewError(400, err.Error()))
			return
		}

		c.JSON(200, web.NewError(200, "The seller has been deleted"))
	}
}
