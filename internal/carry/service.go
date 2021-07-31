package carry

import (
	"context"
	"errors"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
	"github.com/BenjaminBergerM/go-meli-exercise/internal/locality"
)

var ErrUnique = errors.New("cid field has already been taken")
var ErrNotFound = errors.New("cid not exists")
var ErrLocality = errors.New("locality does not exist")

type Service interface {
	Store(ctx context.Context, batch, localty, id int, cid, company, address, telephone string) (domain.Carry, error)
}

type service struct {
	repository Repository
	loc_repo   locality.Repository
}

func NewService(repository Repository, lrepo locality.Repository) Service {
	return &service{
		repository: repository,
		loc_repo:   lrepo,
	}
}

func (s *service) Store(ctx context.Context, batch, localty, id int, cid, company, address, telephone string) (domain.Carry, error) {
	carry := domain.Carry{
		Batch:     batch,
		Locality:  localty,
		Cid:       cid,
		Company:   company,
		Address:   address,
		Telephone: telephone,
		ID:        id,
	}

	if !s.loc_repo.Exists(ctx, localty) {
		return carry, ErrLocality
	}

	if s.repository.Exists(ctx, id) {
		return carry, ErrUnique
	}

	idCarry, err := s.repository.Save(ctx, carry)
	if err != nil {
		return domain.Carry{}, err
	}
	carry.ID = idCarry
	return carry, nil
}
