package handler

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
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

func (e *Employee) Get() gin.HandlerFunc {
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
		employee, err := e.employeeService.Get(ctx, paramID)
		if err != nil {
			c.JSON(404, web.NewError(404, "employee not found"))
			return
		}

		c.JSON(200, &response{employee})
	}
}

func (e *Employee) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Employee `json:"data"`
	}

	return func(c *gin.Context) {

		ctx := context.Background()
		products, err := e.employeeService.GetAll(ctx)
		if err != nil {
			c.JSON(404, web.NewError(404, err.Error()))
			return
		}

		c.JSON(200, &response{products})
	}
}

func (e *Employee) Store() gin.HandlerFunc {
	type request struct {
		CardNumberID string `json:"card_number_id"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		WarehouseID  int    `json:"warehouse_id"`
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
		if req.WarehouseID == 0 {
			c.JSON(422, web.NewError(422, "warehouse_id can not be empty"))
			return
		}

		ctx := context.Background()
		emp, err := e.employeeService.Store(ctx, req.CardNumberID, req.FirstName, req.FirstName, req.WarehouseID)
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

func (e *Employee) Update() gin.HandlerFunc {
	type request struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		WarehouseID int    `json:"warehouse_id"`
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
		emp, err := e.employeeService.Update(ctx, paramID, req.FirstName, req.LastName, req.WarehouseID)
		if err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}

		c.JSON(200, &response{emp})
	}
}

func (e *Employee) Delete() gin.HandlerFunc {
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
		err := e.employeeService.Delete(ctx, paramID)
		if err != nil {
			c.JSON(500, web.NewError(500, err.Error()))
			return
		}

		c.JSON(200, web.NewError(200, "The employee has been deleted"))
	}
}
