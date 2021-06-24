package handler

import (
	"context"
	"strconv"

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
		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.JSON(422, web.NewError(422, "id must be an integer"))
			return
		}

		ctx := context.Background()
		employee, err := e.employeeService.Get(ctx, id)
		if err != nil {

		}

		c.JSON(201, &response{employee})
	}
}

func (e *Employee) GetAll() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
		ctx := context.Background()
		employee, err := e.employeeService.GetAll(ctx)
		if err != nil {

		}
		c.JSON(201, &response{employee})
	}
}

func (e *Employee) Store() gin.HandlerFunc {
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
		employee, err := e.employeeService.Store(ctx, req.FirstName, req.FirstName, req.WarehouseID)
		if err != nil {

		}

		c.JSON(201, &response{employee})
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
		if req.WarehouseID == 0 {
			c.JSON(422, web.NewError(422, "warehouse_id can not be empty"))
			return
		}

		ctx := context.Background()
		employee, err := e.employeeService.Update(ctx, id, req.FirstName, req.FirstName, req.WarehouseID)
		if err != nil {

		}

		c.JSON(201, &response{employee})
	}
}

func (e *Employee) Delete() gin.HandlerFunc {
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
		employee, err := e.employeeService.Delete(ctx, id)
		if err != nil {

		}

		c.JSON(201, &response{employee})
	}
}
