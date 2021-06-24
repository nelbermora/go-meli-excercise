package employee

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

// Service encapsulates the business logic of a Employees.
// As stated by this principle https://golang.org/doc/effective_go#generality,
// since the underlying concrete implementation does not export any other method that is not in the interface,
// we decided to define it where it is implemented rather where it is used (commonly in a handler).
type Service interface {
	Get(ctx context.Context, ID int) (domain.Employee, error)
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Store(ctx context.Context, firstName, lastName string, warehouseID int) (domain.Employee, error)
	Update(ctx context.Context, ID int, firstName, lastName string, warehouseID int) (domain.Employee, error)
	Delete(ctx context.Context, ID int) (domain.Employee, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Get(ctx context.Context, ID int) (domain.Employee, error) {
	e := domain.Employee{}
	// TODO: Implement Repository
	return e, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Employee, error) {
	e := []domain.Employee{}
	// TODO: Implement Repository
	return e, nil
}

func (s *service) Store(ctx context.Context, firstName, lastName string, warehouseID int) (domain.Employee, error) {
	e := domain.Employee{
		FirstName:   firstName,
		LastName:    lastName,
		WarehouseID: warehouseID,
	}
	// TODO: Implement Repository
	return e, nil
}

func (s *service) Update(ctx context.Context, ID int, firstName, lastName string, warehouseID int) (domain.Employee, error) {
	e := domain.Employee{
		FirstName:   firstName,
		LastName:    lastName,
		WarehouseID: warehouseID,
	}
	// TODO: Implement Repository
	return e, nil
}

func (s *service) Delete(ctx context.Context, ID int) (domain.Employee, error) {
	e := domain.Employee{}
	// TODO: Implement Repository
	return e, nil
}
