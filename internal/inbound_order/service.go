package inboundorder

import (
	"context"
	"errors"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/employee"
	"time"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

var (
	ErrEmployeeExistance = errors.New("id employee does not exist")
)

type Service interface {
	Store(ctx context.Context, orderDate time.Time, orderNumber string,employeeId int, productBatchId int, warehouseId int) (domain.InboundOrder, error)
}

type service struct {
	repository Repository
	employeeRepo  employee.Repository
}

func NewService(repository Repository, employeeRepo employee.Repository) Service {
	return &service{
		repository: repository,
		employeeRepo:  employeeRepo,
	}
}

func (s *service) Store(ctx context.Context, orderDate time.Time, orderNumber string,employeeId int, productBatchId int, warehouseId int) (domain.InboundOrder, error) {
	inboundOrder := domain.InboundOrder{
		OrderDate:      orderDate,
		OrderNumber:    orderNumber,
		EmployeeId:     employeeId,
		ProductBatchId: productBatchId,
		WarehouseId:    warehouseId,
	}

	if !s.employeeRepo.ExistsById(ctx, employeeId) {
		return domain.InboundOrder{}, ErrEmployeeExistance
	}

	id, err := s.repository.Save(ctx, inboundOrder)
	if err != nil {
		return domain.InboundOrder{}, err
	}
	inboundOrder.ID = id
	return inboundOrder, nil
}
