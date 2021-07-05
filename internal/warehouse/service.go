package warehouse

import (
	"context"
	"errors"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

var UNIQUE = errors.New("There is a seller with same.")

// Service encapsulates the business logic of a Warehouse.
// As stated by this principle https://golang.org/doc/effective_go#generality,
// since the underlying concrete implementation does not export any other method that is not in the interface,
// we decided to define it where it is implemented rather where it is used (commonly in a handler).
type Service interface {
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Store(ctx context.Context, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature, sectionNumber int) (domain.Warehouse, error)
	Update(ctx context.Context, id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature, sectionNumber int) (domain.Warehouse, error)
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

func (s *service) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature, sectionNumber int) (domain.Warehouse, error) {
	exist := s.repository.Exists(ctx, warehouseCode)
	if exist {
		return domain.Warehouse{}, UNIQUE
	}

	w := domain.Warehouse{
		Address:            address,
		Telephone:          telephone,
		WarehouseCode:      warehouseCode,
		MinimunCapacity:    minimunCapacity,
		MinimunTemperature: minimunTemperature,
		SectionNumber:      sectionNumber,
	}

	id, err := s.repository.Save(ctx, w)
	if err != nil {
		return domain.Warehouse{}, err
	}

	w.ID = id

	return w, nil
}

func (s *service) Update(ctx context.Context, id int, address, telephone, warehouseCode string, minimunCapacity, minimunTemperature, sectionNumber int) (domain.Warehouse, error) {
	w := domain.Warehouse{
		ID:                 id,
		Address:            address,
		Telephone:          telephone,
		WarehouseCode:      warehouseCode,
		MinimunCapacity:    minimunCapacity,
		MinimunTemperature: minimunTemperature,
		SectionNumber:      sectionNumber,
	}

	err := s.repository.Update(ctx, w)
	if err != nil {
		return domain.Warehouse{}, err
	}
	return w, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
