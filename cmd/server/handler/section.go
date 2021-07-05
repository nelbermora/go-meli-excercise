package handler

import (
	"context"
	"strconv"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/section"

	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type Section struct {
	sectionService section.Service
}

func NewSection(s section.Service) *Section {
	return &Section{
		sectionService: s,
	}
}

func (s *Section) GetAll() gin.HandlerFunc {

	type response struct {
		Data []domain.Section `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		sections, err := s.sectionService.GetAll(ctx)
		if err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		}

		c.JSON(200, &response{sections})
	}
}

func (s *Section) Get() gin.HandlerFunc {

	type response struct {
		Data domain.Section `json:"data"`
	}

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid ID"))
			return
		}
		ctx := context.Background()
		sel, err := s.sectionService.Get(ctx, int(id))
		if err != nil {
			c.JSON(404, web.NewError(404, "section not found"))
			return
		}
		c.JSON(201, &response{sel})
	}
}

func (s *Section) Store() gin.HandlerFunc {
	type request struct {
		SectionNumber      int `json:"section_number"`
		CurrentTemperature int `json:"current_temperature"`
		MinTemperature     int `json:"minimum_temperature"`
		CurrentCapacity    int `json:"current_capacity"`
		MinCapacity        int `json:"minimum_capacity"`
		MaxCapacity        int `json:"maximum_capacity"`
		WarehouseID        int `json:"warehouse_id"`
		ProductTypeID      int `json:"product_type_id"`
	}

	type response struct {
		Data domain.Section `json:"data"`
	}

	return func(c *gin.Context) {
		var req request

		if err := c.Bind(&req); err != nil {
			c.JSON(422, web.NewError(400, "json decoding: "+err.Error()))
			return
		}
		if req.SectionNumber == 0 {
			c.JSON(422, web.NewError(422, "section_number can not be empty"))
			return
		}
		if req.CurrentTemperature == 0 {
			c.JSON(422, web.NewError(422, "current_temperature can not be empty"))
			return
		}
		if req.MinTemperature == 0 {
			c.JSON(422, web.NewError(422, "minimum_temperature can not be empty"))
			return
		}
		if req.CurrentCapacity == 0 {
			c.JSON(422, web.NewError(422, "current_capacity can not be empty"))
			return
		}
		if req.MinCapacity == 0 {
			c.JSON(422, web.NewError(422, "minimum_capacity can not be empty"))
			return
		}
		if req.MaxCapacity == 0 {
			c.JSON(422, web.NewError(422, "maximum_capacity can not be empty"))
			return
		}
		if req.WarehouseID == 0 {
			c.JSON(422, web.NewError(422, "warehouse_id can not be empty"))
			return
		}
		if req.ProductTypeID == 0 {
			c.JSON(422, web.NewError(422, "product_type_id can not be empty"))
			return
		}
		ctx := context.Background()
		sec, err := s.sectionService.Store(ctx, req.SectionNumber, req.CurrentTemperature, req.MinTemperature, req.CurrentCapacity, req.MinCapacity, req.MaxCapacity, req.WarehouseID, req.ProductTypeID)
		if err != nil {
			switch err {
			case section.UNIQUE:
				c.JSON(409, web.NewError(409, err.Error()))
			default:
				c.JSON(500, web.NewError(500, err.Error()))
			}
			return
		}

		c.JSON(201, &response{sec})
	}
}

func (s *Section) Update() gin.HandlerFunc {

	type request struct {
		SectionNumber      int `json:"section_number"`
		CurrentTemperature int `json:"current_temperature"`
		MinTemperature     int `json:"minimum_temperature"`
		CurrentCapacity    int `json:"current_capacity"`
		MinCapacity        int `json:"minimum_capacity"`
		MaxCapacity        int `json:"maximum_capacity"`
		WarehouseID        int `json:"warehouse_id"`
		ProductTypeID      int `json:"product_type_id"`
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
		sel, err := s.sectionService.Update(ctx, int(id), req.SectionNumber, req.CurrentTemperature, req.MinTemperature, req.CurrentCapacity, req.MinCapacity, req.MaxCapacity, req.WarehouseID, req.ProductTypeID)
		if err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}

		c.JSON(200, sel)
	}
}

func (s *Section) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, web.NewError(400, "invalid ID"))
			return
		}

		ctx := context.Background()
		err = s.sectionService.Delete(ctx, int(id))
		if err != nil {
			c.JSON(400, web.NewError(400, err.Error()))
			return
		}

		c.JSON(200, web.NewError(200, "The section has been deleted"))
	}
}
