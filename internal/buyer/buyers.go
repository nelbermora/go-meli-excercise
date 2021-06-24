package buyer

import (
	"context"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

// Service encapsulates the business logic of a Warehouse.
// As stated by this principle https://golang.org/doc/effective_go#generality,
// since the underlying concrete implementation does not export any other method that is not in the interface,
// we decided to define it where it is implemented rather where it is used (commonly in a handler).
type Service interface {
	Get(ctx context.Context, ID int) (domain.Buyer, error)
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Store(ctx context.Context, firstName, lastName string) (domain.Buyer, error)
	Update(ctx context.Context, ID int, firstName, lastName string) (domain.Buyer, error)
	Delete(ctx context.Context, ID int) (domain.Buyer, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Get(ctx context.Context, ID int) (domain.Buyer, error) {
	b := domain.Buyer{}
	// TODO: Implement Repository
	return b, nil
}

func (s *service) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	b := []domain.Buyer{}
	// TODO: Implement Repository
	return b, nil
}

func (s *service) Store(ctx context.Context, firstName, lastName string) (domain.Buyer, error) {
	b := domain.Buyer{
		FirstName: firstName,
		LastName:  lastName,
	}

	return b, nil
}

func (s *service) Update(ctx context.Context, ID int, firstName, lastName string) (domain.Buyer, error) {
	b := domain.Buyer{
		FirstName: firstName,
		LastName:  lastName,
	}
	// TODO: Implement Repository
	return b, nil
}

func (s *service) Delete(ctx context.Context, ID int) (domain.Buyer, error) {
	b := domain.Buyer{}
	// TODO: Implement Repository
	return b, nil
}
