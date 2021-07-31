package productbatch

import (
	"context"
	"errors"
	"time"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/product"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/section"
)

var (
	ErrUnique           = errors.New("locality_id field has already been taken")
	ErrNotFound         = errors.New("locality_id not exists")
	ErrProdExistance    = errors.New("id product does not exist")
	ErrSectionExistance = errors.New("id section does not exist")
)

type Service interface {
	Store(ctx context.Context, qty, initQty, prod, section int, temp, minTemp float32, dueDate, manufDate time.Time, batch string) (domain.ProdcutBatch, error)
}

type service struct {
	repository  Repository
	sectionRepo section.Repository
	prodRepo    product.Repository
}

func NewService(repository Repository, sectionRepo section.Repository, prodRepo product.Repository) Service {
	return &service{
		repository:  repository,
		sectionRepo: sectionRepo,
		prodRepo:    prodRepo,
	}
}

func (s *service) Store(ctx context.Context, qty, initQty, prod, section int, temp, minTemp float32, dueDate, manufDate time.Time, batch string) (domain.ProdcutBatch, error) {
	productBatch := domain.ProdcutBatch{
		BatchNumber:       batch,
		CurrentQuantity:   qty,
		CurrentTemp:       temp,
		DueDate:           dueDate,
		InitialQuantity:   initQty,
		ManufacturingDate: manufDate,
		MinTemperature:    minTemp,
		ProductId:         prod,
		SectionId:         section,
	}

	if !s.prodRepo.ExistsById(ctx, prod) {
		return domain.ProdcutBatch{}, ErrProdExistance
	}

	if !s.sectionRepo.ExistsById(ctx, section) {
		return domain.ProdcutBatch{}, ErrSectionExistance
	}

	idProduct, err := s.repository.Save(ctx, productBatch)
	if err != nil {
		return domain.ProdcutBatch{}, err
	}
	productBatch.ID = idProduct
	return productBatch, nil
}
