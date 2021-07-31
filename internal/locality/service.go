package locality

import (
	"context"
	"errors"

	"github.com/BenjaminBergerM/go-meli-exercise/internal/domain"
)

var ErrUnique = errors.New("locality_id field has already been taken")
var ErrNotFound = errors.New("locality_id not exists")

type Service interface {
	Store(ctx context.Context, id int, name, province, country string) (domain.Locality, error)
	GetSellersByLoc(ctx context.Context, id int) (domain.Locality, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Store(ctx context.Context, id int, name, province, country string) (domain.Locality, error) {
	locality := domain.Locality{
		ID:       id,
		Name:     name,
		Province: province,
		Country:  country,
	}

	if s.repository.Exists(ctx, id) {
		return locality, ErrUnique
	}

	idLocality, err := s.repository.Save(ctx, locality)
	if err != nil {
		return domain.Locality{}, err
	}
	locality.ID = idLocality
	return locality, nil
}

func (s *service) GetSellersByLoc(ctx context.Context, id int) (domain.Locality, error) {
	// checking if locality exists
	if !s.repository.Exists(ctx, id) {
		return domain.Locality{}, ErrNotFound
	}
	l, err := s.repository.GetSellersByLoc(ctx, id)
	if err != nil {
		// if locality exists, but repository returns nothing, so locality has 0 sellers
		// looking for locality info
		l, err = s.repository.Get(ctx, id)
		if err != nil {
			return domain.Locality{}, err
		}
		// setting seller_count = 0
		count := 0
		l.Sellers = &count
		return l, nil
	}
	return l, nil
}
