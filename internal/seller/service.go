package seller

import (
	"context"
	"errors"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

var UNIQUE = errors.New("There is a seller with same.")

// Service encapsulates the business logic of a Seller.
// As stated by this principle https://golang.org/doc/effective_go#generality,
// since the underlying concrete implementation does not export any other method that is not in the interface,
// we decided to define it where it is implemented rather where it is used (commonly in a handler).
type Service interface {
	Get(ctx context.Context, id int) (domain.Seller, error)
	GetAll(ctx context.Context) ([]domain.Seller, error)
	Store(ctx context.Context, cid int, companyName, address, telephone string, localityID int) (domain.Seller, error)
	Update(ctx context.Context, id, cid int, companyName, address, telephone string, localityID int) (domain.Seller, error)
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

func (s *service) Get(ctx context.Context, id int) (domain.Seller, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]domain.Seller, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, cid int, companyName, address, telephone string, localityID int) (domain.Seller, error) {

	exist := s.repository.Exists(ctx, cid)
	if exist {
		return domain.Seller{}, UNIQUE
	}

	seller := domain.Seller{
		CID:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
		LocalityID:  localityID,
	}

	id, err := s.repository.Save(ctx, seller)
	if err != nil {
		return domain.Seller{}, err
	}

	seller.ID = id

	return seller, nil
}

func (s *service) Update(ctx context.Context, id, cid int, companyName, address, telephone string, localityID int) (domain.Seller, error) {
	seller := domain.Seller{
		ID:          id,
		CID:         cid,
		CompanyName: companyName,
		Address:     address,
		Telephone:   telephone,
		LocalityID:  localityID,
	}

	err := s.repository.Update(ctx, seller)
	if err != nil {
		return domain.Seller{}, err
	}

	return seller, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
