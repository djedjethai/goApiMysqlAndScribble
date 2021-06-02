package listing

import (
	"errors"
)

var ErrNotFound = errors.New("Beer not found")

type Service interface {
	GetBeerS(string) (Beer, error)
	GetAllBeersS() ([]Beer, error)
	GetBeerReviewS(string) ([]Review, error)
}

type Repository interface {
	GetBeer(string) (Beer, error)
	GetAllBeers() ([]Beer, error)
	GetBeerReviews(string) ([]Review, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetBeerReviewS(bid string) ([]Review, error) {
	return s.r.GetBeerReviews(bid)
}

func (s *service) GetAllBeersS() ([]Beer, error) {
	return s.r.GetAllBeers()
}

func (s *service) GetBeerS(bid string) (Beer, error) {
	return s.r.GetBeer(bid)
}
