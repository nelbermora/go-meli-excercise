package product

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

// Service encapsulates the business logic of a Product.
// As stated by this principle https://golang.org/doc/effective_go#generality,
// since the underlying concrete implementation does not export any other method that is not in the interface,
// we decided to define it where it is implemented rather where it is used (commonly in a handler).
type Service interface {
	Get(ctx context.Context, id int) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Store(ctx context.Context, description, productCode string, height, length, netweight, recomFreezTemp, width float32, productTypeID, sellerID, expirationRate, freezingRate int) (domain.Product, error)
	Update(ctx context.Context, id int, description, productCode string, height, length, netweight, recomFreezTemp, width float32, productTypeID, sellerID, expirationRate, freezingRate int) (domain.Product, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Get(ctx context.Context, id int) (domain.Product, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, description, productCode string, height, length, netweight, recomFreezTemp, width float32, productTypeID, sellerID, expirationRate, freezingRate int) (domain.Product, error) {
	p := domain.Product{
		Description:    description,
		ExpirationRate: expirationRate,
		FreezingRate:   freezingRate,
		Height:         height,
		Length:         length,
		Netweight:      netweight,
		ProductCode:    productCode,
		RecomFreezTemp: recomFreezTemp,
		Width:          width,
		ProductTypeID:  productTypeID,
		SellerID:       sellerID,
	}

	id, err := s.repository.Save(ctx, p)
	if err != nil {
		return domain.Product{}, err
	}

	p.ID = id

	return p, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *service) Update(ctx context.Context, id int, description, productCode string, height, length, netweight, recomFreezTemp, width float32, productTypeID, sellerID, expirationRate, freezingRate int) (domain.Product, error) {
	p := domain.Product{
		ID:             id,
		Description:    description,
		ExpirationRate: expirationRate,
		FreezingRate:   freezingRate,
		Height:         height,
		Length:         length,
		Netweight:      netweight,
		ProductCode:    productCode,
		RecomFreezTemp: recomFreezTemp,
		Width:          width,
		ProductTypeID:  productTypeID,
		SellerID:       sellerID,
	}

	err := s.repository.Update(ctx, p)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}
