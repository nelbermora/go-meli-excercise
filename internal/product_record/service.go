package productrecord

import (
	"context"
	"errors"
	"time"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/product"
)

var (
	ErrProdExistance = errors.New("id product does not exist")
)

type Service interface {
	Store(ctx context.Context, update time.Time, purchase, sale float32, product int) (domain.ProductRecord, error)
}

type service struct {
	repository Repository
	prodRepo   product.Repository
}

func NewService(repository Repository, prodRepo product.Repository) Service {
	return &service{
		repository: repository,
		prodRepo:   prodRepo,
	}
}

func (s *service) Store(ctx context.Context, update time.Time, purchase, sale float32, product int) (domain.ProductRecord, error) {
	productRecord := domain.ProductRecord{
		LastUpdate:    update,
		PurchasePrice: purchase,
		SalePrice:     sale,
		ProductId:     product,
	}

	if !s.prodRepo.ExistsById(ctx, product) {
		return domain.ProductRecord{}, ErrProdExistance
	}

	idProduct, err := s.repository.Save(ctx, productRecord)
	if err != nil {
		return domain.ProductRecord{}, err
	}
	productRecord.ID = idProduct
	return productRecord, nil
}
