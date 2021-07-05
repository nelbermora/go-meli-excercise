package section

import (
	"context"
	"errors"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

var UNIQUE = errors.New("The section_number field has already been taken.")

// Service encapsulates the business logic of a Section.
// As stated by this principle https://golang.org/doc/effective_go#generality,
// since the underlying concrete implementation does not export any other method that is not in the interface,
// we decided to define it where it is implemented rather where it is used (commonly in a handler).
type Service interface {
	Get(ctx context.Context, id int) (domain.Section, error)
	GetAll(ctx context.Context) ([]domain.Section, error)
	Store(ctx context.Context, sectionNumber, currentTemperature, minTemperature, currentCapacity, minCapacity, maxCapacity, warehouseID, ProductTypeID int) (domain.Section, error)
	Update(ctx context.Context, id, sectionNumber, currentTemperature, minTemperature, currentCapacity, minCapacity, maxCapacity, warehouseID, productTypeID int) (domain.Section, error)
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

func (s *service) Get(ctx context.Context, id int) (domain.Section, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Section, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, sectionNumber, currentTemperature, minTemperature, currentCapacity, minCapacity, maxCapacity, warehouseID, productTypeID int) (domain.Section, error) {

	exist := s.repository.Exists(ctx, sectionNumber)
	if exist {
		return domain.Section{}, UNIQUE
	}

	section := domain.Section{
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minCapacity,
		MaximumCapacity:    maxCapacity,
		WarehouseID:        warehouseID,
		ProductTypeID:      productTypeID,
	}

	id, err := s.repository.Save(ctx, section)
	if err != nil {
		return domain.Section{}, err
	}

	section.ID = id

	return section, nil
}

func (s *service) Update(ctx context.Context, id, sectionNumber, currentTemperature, minTemperature, currentCapacity, minCapacity, maxCapacity, warehouseID, productTypeID int) (domain.Section, error) {
	section := domain.Section{
		ID:                 id,
		SectionNumber:      sectionNumber,
		CurrentTemperature: currentTemperature,
		MinimumTemperature: minTemperature,
		CurrentCapacity:    currentCapacity,
		MinimumCapacity:    minCapacity,
		MaximumCapacity:    maxCapacity,
		WarehouseID:        warehouseID,
		ProductTypeID:      productTypeID,
	}

	err := s.repository.Update(ctx, section)
	if err != nil {
		return domain.Section{}, err
	}

	return section, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
