package handler

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/employee"
	"github.com/BenjaminBergerM/go-meli-exercise/pkg/web"
	"github.com/gin-gonic/gin"
)

type Employee struct {
	employeeService employee.Service
}

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		employeeService: e,
	}
}

func (w *Employee) Get() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}

func (w *Employee) Store() gin.HandlerFunc {
	type request struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		WarehouseID int    `json:"warehouse_id"`
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
		if req.WarehouseID == 0 {
			c.JSON(422, web.NewError(422, "warehouse_id can not be empty"))
			return
		}

		ctx := context.Background()
		warehouse, err := w.employeeService.Store(ctx, req.FirstName, req.FirstName, req.WarehouseID)
		if err != nil {

		}

		c.JSON(201, &response{warehouse})
	}
}

func (w *Employee) Update() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}

func (w *Employee) Delete() gin.HandlerFunc {
	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

	}
}
