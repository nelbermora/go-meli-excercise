package purchaseorder

import (
	"context"
	"errors"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/buyer"
	"time"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

var (
	ErrBuyerExistance = errors.New("id buyer does not exist")
)

type Service interface {
	Store(ctx context.Context, orderNumber string, orderDate time.Time, trackingCode string, buyerId int, productRecordId int, orderStatusId int) (domain.PurchaseOrder, error)
}

type service struct {
	repository Repository
	buyerRepo  buyer.Repository
}

func NewService(repository Repository, prodRepo buyer.Repository) Service {
	return &service{
		repository: repository,
		buyerRepo:  prodRepo,
	}
}

func (s *service) Store(ctx context.Context, orderNumber string, orderDate time.Time, trackingCode string, buyerId int, productRecordId int, orderStatusId int) (domain.PurchaseOrder, error) {
	purchaseOrder := domain.PurchaseOrder{
		OrderNumber:     orderNumber,
		OrderDate:       orderDate,
		TrackingCode:    trackingCode,
		BuyerId:         buyerId,
		ProductRecordId: productRecordId,
		OrderStatusId:   orderStatusId,
	}

	if !s.buyerRepo.ExistsById(ctx, buyerId) {
		return domain.PurchaseOrder{}, ErrBuyerExistance
	}

	id, err := s.repository.Save(ctx, purchaseOrder)
	if err != nil {
		return domain.PurchaseOrder{}, err
	}
	purchaseOrder.ID = id
	return purchaseOrder, nil
}
