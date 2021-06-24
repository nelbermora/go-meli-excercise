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
	GetAll(ctx context.Context) ([]domain.Product, error)
	Store(ctx context.Context, description, productCode string, height, length, netweight, recomFreezTemp, width float32, productTypeID, sellerID, expirationRate, freezingRate int) (domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Store(ctx context.Context, description, productCode string, height, length, netweight, recomFreezTemp, width float32, productTypeID, sellerID, expirationRate, freezingRate int) (domain.Product, error) {
	p := domain.Product{
		Description: description,
		ExpirationRate: expirationRate,
		FreezingRate: freezingRate,
		Height: height,
		Length: length,
		Netweight: netweight,
		ProductCode: productCode,
		RecomFreezTemp: recomFreezTemp,
		Width: width,
		ProductTypeID: productTypeID,
		SellerID: sellerID,
	}

	id, err := s.repository.Save(ctx, p)
	if err != nil {
		return domain.Product{}, err
	}

	p.ID = id

	return p, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repository.GetAll(ctx)
}