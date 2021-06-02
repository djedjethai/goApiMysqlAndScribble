package reviewing

import (
	"errors"
)

var ErrNotFound = errors.New("Reviews not found")
var ErrRegister = errors.New("Server error")

type Service interface {
	// AddBeerReviewS(Review) (string, error)
	AddReviewSampleS([]Review) error
	PostBeerReviewS(Review) (string, error)
}

type Repository interface {
	AddBeerReview(Review) (string, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) PostBeerReviewS(rev Review) (string, error) {
	return s.r.AddBeerReview(rev)
}

func (s *service) AddReviewSampleS(rvs []Review) error {
	for _, rv := range rvs {
		_, err := s.r.AddBeerReview(rv)
		if err != nil {
			return err
		}
	}

	return nil
}

// func (s *service) AddBeerReviewS(rv Review) (string, error) {
// 	id, err := s.r.AddBeerReview(rv)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	return id, nil
// }
