package handler

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/carry"
	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type Carry struct {
	carryService carry.Service
}

func NewCarry(c carry.Service) *Carry {
	return &Carry{
		carryService: c,
	}
}

func (ca *Carry) Store() gin.HandlerFunc {
	type request struct {
		ID        int    `json:"id"`
		Cid       string `json:"cid"`
		Batch     int    `json:"batch_number"`
		Company   string `json:"company_name"`
		Address   string `json:"adress"`
		Telephone string `json:"telephon"`
		Locality  int    `json:"locality_id"`
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
		if req.ID == 0 {
			c.JSON(422, web.NewError(422, "id must be greater than 0"))
			return
		}
		if req.Cid == "" {
			c.JSON(422, web.NewError(422, "CID can not be empty"))
			return
		}
		if req.Batch == 0 {
			c.JSON(422, web.NewError(422, "Batch Number can not be empty"))
			return
		}
		if req.Company == "" {
			c.JSON(422, web.NewError(422, "company can not be empty"))
			return
		}
		if req.Locality == 0 {
			c.JSON(422, web.NewError(422, "locality_id must be greather than 0"))
			return
		}

		ctx := context.Background()
		result, err := ca.carryService.Store(ctx, req.Batch, req.Locality, req.ID, req.Cid, req.Company, req.Address, req.Telephone)
		if err != nil {
			switch err {
			case carry.ErrNotFound:
				c.JSON(409, web.NewError(409, err.Error()))
			case carry.ErrLocality:
				c.JSON(409, web.NewError(409, err.Error()))
			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}
		c.JSON(201, &response{result})
	}
}
