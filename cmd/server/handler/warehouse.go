package handler

import (
	"github.com/BenjaminBergerM/go-meli-exercise/internal/warehouse"
	"github.com/gin-gonic/gin"
)

type Warehouse struct {
	warehouseService warehouse.Service
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		warehouseService: w,
	}
}

func (w *Warehouse) Get() gin.HandlerFunc {
	type Request struct {
	}

	type Response struct {
	}

	return func(c *gin.Context) {

	}
}
