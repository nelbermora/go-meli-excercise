package employee

import (
	"context"
	"errors"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

var UNIQUE = errors.New("The card_number_id field has already been taken.")
var ErrNotFound = errors.New("employee_id not exists")


// Service encapsulates the business logic of a employee.
// As stated by this principle https://golang.org/doc/effective_go#generality,
// since the underlying concrete implementation does not export any other method that is not in the interface,
// we decided to define it where it is implemented rather where it is used (commonly in a handler).
type Service interface {
	Get(ctx context.Context, cardNumberID string) (domain.Employee, error)
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Store(ctx context.Context, cardNumberID, firstName, lastName string, warehouseID int) (domain.Employee, error)
	Update(ctx context.Context, cardNumberID, firstName, lastName string, warehouseID int) (domain.Employee, error)
	Delete(ctx context.Context, cardNumberID string) error
	GetInboundOrdersByEmployee(ctx context.Context, id int) ([]domain.Employee, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Get(ctx context.Context, cardNumberID string) (domain.Employee, error) {
	return s.repository.Get(ctx, cardNumberID)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Employee, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, cardNumberID, firstName, lastName string, warehouseID int) (domain.Employee, error) {

	exist := s.repository.Exists(ctx, cardNumberID)
	if exist {
		return domain.Employee{}, UNIQUE
	}

	employee := domain.Employee{
		CardNumberID: cardNumberID,
		FirstName:    firstName,
		LastName:     lastName,
		WarehouseID:  warehouseID,
	}

	id, err := s.repository.Save(ctx, employee)
	if err != nil {
		return domain.Employee{}, err
	}

	employee.ID = id

	return employee, nil
}

func (s *service) Update(ctx context.Context, cardNumberID, firstName, lastName string, warehouseID int) (domain.Employee, error) {

	employee := domain.Employee{
		CardNumberID: cardNumberID,
		FirstName:    firstName,
		LastName:     lastName,
		WarehouseID:  warehouseID,
	}

	err := s.repository.Update(ctx, employee)
	if err != nil {
		return domain.Employee{}, err
	}

	return employee, nil
}

func (s *service) Delete(ctx context.Context, cardNumberID string) error {
	return s.repository.Delete(ctx, cardNumberID)
}

func (s *service) GetInboundOrdersByEmployee(ctx context.Context, id int) ([]domain.Employee, error) {
	result := []domain.Employee{}
	if !s.repository.ExistsById(ctx, id) && id > 0 {
		return result, ErrNotFound
	}

	result, err := s.repository.GetInboundOrdersByEmployee(ctx, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
