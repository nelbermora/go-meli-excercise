package handler

import (
	"context"
	"strconv"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/locality"
	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type Locality struct {
	localotyService locality.Service
}

func NewLocality(l locality.Service) *Locality {
	return &Locality{
		localotyService: l,
	}
}

func (l *Locality) Store() gin.HandlerFunc {
	type request struct {
		ID       int    `json:"id"`
		Name     string `json:"locality_name"`
		Province string `json:"province_name"`
		Country  string `json:"country_name"`
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
		if req.Name == "" {
			c.JSON(422, web.NewError(422, "locality_name can not be empty"))
			return
		}
		if req.Province == "" {
			c.JSON(422, web.NewError(422, "province_name can not be empty"))
			return
		}
		if req.Country == "" {
			c.JSON(422, web.NewError(422, "country_name can not be empty"))
			return
		}

		ctx := context.Background()
		emp, err := l.localotyService.Store(ctx, req.ID, req.Name, req.Province, req.Country)
		if err != nil {
			switch err {
			case locality.ErrUnique:
				c.JSON(409, web.NewError(409, err.Error()))
			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}
		c.JSON(201, &response{emp})
	}
}

func (l *Locality) GetSellersByLoc() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			c.JSON(422, web.NewError(422, "id must be distinct of null"))
			return
		}
		intId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(422, web.NewError(422, "invalid Id"))
			return
		}

		ctx := context.Background()
		rep, err := l.localotyService.GetSellersByLoc(ctx, intId)
		if err != nil {
			c.JSON(404, web.NewError(404, "locality not found"))
			return
		}
		c.JSON(200, &response{rep})
	}
}
