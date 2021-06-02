package updating

import (
	"errors"
	"github.com/djedjethai/mytry/pkg/listing"
)

var ErrRegister = errors.New("err during updating")

type Service interface {
	UpdateReviewS(listing.Review) (string, error)
	UpdateBeerS(listing.Beer) (string, error)
}

type service struct {
	r Repository
}

type Repository interface {
	UpdateReview(listing.Review) (string, error)
	UpdateBeer(listing.Beer) (string, error)
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) UpdateReviewS(review listing.Review) (string, error) {
	return s.r.UpdateReview(review)
}

func (s *service) UpdateBeerS(beerUpdate listing.Beer) (string, error) {
	return s.r.UpdateBeer(beerUpdate)
}
