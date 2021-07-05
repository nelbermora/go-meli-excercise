package buyer

import (
	"context"
	"errors"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

var UNIQUE = errors.New("The card_number_id field has already been taken.")

// Service encapsulates the business logic of a buyer.
// As stated by this principle https://golang.org/doc/effective_go#generality,
// since the underlying concrete implementation does not export any other method that is not in the interface,
// we decided to define it where it is implemented rather where it is used (commonly in a handler).
type Service interface {
	Get(ctx context.Context, cardNumberID string) (domain.Buyer, error)
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Store(ctx context.Context, cardNumberID, firstName, lastName string) (domain.Buyer, error)
	Update(ctx context.Context, cardNumberID, firstName, lastName string) (domain.Buyer, error)
	Delete(ctx context.Context, cardNumberID string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Get(ctx context.Context, cardNumberID string) (domain.Buyer, error) {
	return s.repository.Get(ctx, cardNumberID)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Buyer, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, cardNumberID, firstName, lastName string) (domain.Buyer, error) {

	exist := s.repository.Exists(ctx, cardNumberID)
	if exist {
		return domain.Buyer{}, UNIQUE
	}

	buyer := domain.Buyer{
		CardNumberID: cardNumberID,
		FirstName:    firstName,
		LastName:     lastName,
	}

	id, err := s.repository.Save(ctx, buyer)
	if err != nil {
		return domain.Buyer{}, err
	}

	buyer.ID = id

	return buyer, nil
}

func (s *service) Update(ctx context.Context, cardNumberID, firstName, lastName string) (domain.Buyer, error) {

	buyer := domain.Buyer{
		CardNumberID: cardNumberID,
		FirstName:    firstName,
		LastName:     lastName,
	}

	err := s.repository.Update(ctx, buyer)
	if err != nil {
		return domain.Buyer{}, err
	}

	return buyer, nil
}

func (s *service) Delete(ctx context.Context, cardNumberID string) error {
	return s.repository.Delete(ctx, cardNumberID)
}
