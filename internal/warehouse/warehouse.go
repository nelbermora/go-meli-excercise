package warehouse

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

// Service encapsulates the business logic of a Warehouse.
// As stated by this principle https://golang.org/doc/effective_go#generality,
// since the underlying concrete implementation does not export any other method that is not in the interface,
// we decided to define it where it is implemented rather where it is used (commonly in a handler).
type Service interface {
	Get(ctx context.Context, ID int) (domain.Warehouse, error)
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Store(ctx context.Context, address, telephone, warehouseCode string) (domain.Warehouse, error)
	Update(ctx context.Context, ID int, address, telephone, warehouseCode string) (domain.Warehouse, error)
	Delete(ctx context.Context, ID int) (domain.Warehouse, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Get(ctx context.Context, ID int) (domain.Warehouse, error) {
	e := domain.Warehouse{}
	// TODO: Implement Repository
	return e, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	e := []domain.Warehouse{}
	// TODO: Implement Repository
	return e, nil
}

func (s *service) Store(ctx context.Context, address, telephone, warehouseCode string) (domain.Warehouse, error) {
	w := domain.Warehouse{
		Address:       address,
		Telephone:     telephone,
		WarehouseCode: warehouseCode,
	}

	return w, nil
}

func (s *service) Update(ctx context.Context, ID int, address, telephone, warehouseCode string) (domain.Warehouse, error) {
	e := domain.Warehouse{}
	// TODO: Implement Repository
	return e, nil
}

func (s *service) Delete(ctx context.Context, ID int) (domain.Warehouse, error) {
	e := domain.Warehouse{}
	// TODO: Implement Repository
	return e, nil
}
