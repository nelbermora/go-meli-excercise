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
	GetSellersByLoc(ctx context.Context, id int) ([]domain.Locality, error)
	GetCarriesByLoc(ctx context.Context, id int) ([]domain.Locality, error)
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

func (s *service) GetCarriesByLoc(ctx context.Context, id int) ([]domain.Locality, error) {
	// checking if locality exists
	result := []domain.Locality{}
	if !s.repository.Exists(ctx, id) && id > 0 {
		return result, ErrNotFound
	}

	result, err := s.repository.GetCarriesByLoc(ctx, id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *service) GetSellersByLoc(ctx context.Context, id int) ([]domain.Locality, error) {
	// checking if locality exists
	result := []domain.Locality{}
	if !s.repository.Exists(ctx, id) && id > 0 {
		return result, ErrNotFound
	}

	result, err := s.repository.GetSellersByLoc(ctx, id)
	if err != nil {
		return result, err
	}
	return result, nil
}
