package adding

import (
	"errors"
	"github.com/djedjethai/mytry/pkg/listing"
)

var ErrDuplicate = errors.New("beer already exist")
var ErrRegister = errors.New("Server Error")

type Service interface {
	AddBeerS(Beer) (string, error)
	AddBeerSampleS([]Beer) error
}

type Repository interface {
	AddBeer(Beer) (string, error)
	GetAllBeers() ([]listing.Beer, error)
}

type RepoDatabase interface {
	AddBeerDB(Beer) (string, error)
}

type service struct {
	r   Repository
	rdb RepoDatabase
}

func NewService(r Repository, rdb RepoDatabase) Service {
	return &service{r, rdb}
}

func (s *service) AddBeerS(b Beer) (string, error) {
	var ab []listing.Beer

	ab, err := s.r.GetAllBeers()
	if err != nil {
		return "", ErrRegister
	}

	for i := range ab {
		if b.Name == ab[i].Name &&
			b.Brewery == ab[i].Brewery &&
			b.Abv == ab[i].Abv &&
			b.ShortDesc == ab[i].ShortDesc {
			return "", ErrDuplicate
		}
	}

	// add beer in cache
	bid, err := s.r.AddBeer(b)
	if err != nil {
		return "", ErrRegister
	}

	// add beer in db
	_, _ = s.rdb.AddBeerDB(b)

	return bid, nil
}

func (s *service) AddBeerSampleS(beers []Beer) error {
	for i := range beers {
		_, err := s.r.AddBeer(beers[i])
		if err != nil {
			return ErrRegister
		}
	}

	return nil
}
